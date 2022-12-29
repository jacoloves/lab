
#include "ttyclockcopy.h"

void init(void)
{
    struct sigaction sig;
    setlocale(LC_TIME, "");

    ttyclock.bg = COLOR_BLACK;

    // init ncurses
    if (ttyclock.tty) {
        FILE *ftty = fopen(ttyclock.tty, "r+");
        if (!ftty) {
            fprintf(stderr, "tty-clock: error: '%s' couldn't be opend: %s.\n",
                    ttyclock.tty, strerror(errno));
            exit(EXIT_FAILURE);
        }
        ttyclock.ttyscr = newterm(NULL, ftty, ftty);
        assert(ttyclock.ttyscr != NULL);
        set_term(ttyclock.ttyscr);
    } else {
        initscr();
    }

    cbreak();
    noecho();
    keypad(stdscr, TRUE);
    start_color();
    curs_set(false);
    clear();

    // init default terminal color
    if(use_default_colors() == OK)
        ttyclock.bg = -1;

    // init color pair
    init_pair(0, ttyclock.bg, ttyclock.bg);
    init_pair(1, ttyclock.bg, ttyclock.option.color);
    init_pair(2, ttyclock.option.color, ttyclock.bg);

    refresh();

    // init signal handler
    sig.sa_handler = signal_handler;
    sig.sa_flags = 0;
    sigaction(SIGTERM, &sig, NULL);
    sigaction(SIGINT, &sig, NULL);
    sigaction(SIGSEGV, &sig, NULL);

    // Init global struct
    ttyclock.running = true;
    if(!ttyclock.geo.x)
        ttyclock.geo.x = 0;
    if(!ttyclock.geo.y)
        ttyclock.geo.y = 0;
    if(!ttyclock.geo.a)
        ttyclock.geo.a = 1;
    if(!ttyclock.geo.b)
        ttyclock.geo.b = 1;
    ttyclock.geo.w = (ttyclock.option.second) ? SECFRAMEW : NORMFRAMEW;
    ttyclock.geo.h = 7;
    ttyclock.tm = localtime(&(ttyclock.lt));
    if(ttyclock.option.utc) {
        ttyclock.tm = gmtime(&(ttyclock.lt));
    }
    ttyclock.lt = time(NULL);
    update_hour();

    /* Create clock win */
    ttyclock.framewin = newwin(ttyclock.geo.h,
                               ttyclock.geo.w,
                               ttyclock.geo.x,
                               ttyclock.geo.y);
    if(ttyclock.option.box) {
        box(ttyclock.framewin, 0, 0);
    }

    if(ttyclock.option.bold) {
        wattron(ttyclock.framewin, A_BLINK);
    }

    /* Create the date win */
    ttyclock.datewin = newwin(DATEWINH, strlen(ttyclock.date.datestr) + 2,
                              ttyclock.geo.x + ttyclock.geo.h -1,
                              ttyclock.geo.y + (ttyclock.geo.w / 2) -
                              (strlen(ttyclock.date.datestr) / 2) - 1);
    if(ttyclock.option.box && ttyclock.option.date) {
        box(ttyclock.datewin, 0, 0);
    }
    clearok(ttyclock.datewin, true);

    set_center(ttyclock.option.center);

    nodelay(stdscr, true);

    if(ttyclock.option.date) {
        wrefresh(ttyclock.datewin);
    }

    wrefresh(ttyclock.framewin);

    return;
}

void signal_handler(int signal)
{
    switch (signal) {
        case SIGINT:
        case SIGTERM:
            ttyclock.running = false;
            break;
            // Segmentation fault signal
        case SIGSEGV:
            endwin();
            fprintf(stderr, "Segmentation fault.\n");
            exit(EXIT_FAILURE);
    }

    return;
}

void cleanup(void)
{
    if (ttyclock.ttyscr)
        delscreen(ttyclock.ttyscr);

    free(ttyclock.tty);
}

void clock_move(int x, int y, int w, int h)
{
    /* Erase border for a clean move */
    wbkgdset(ttyclock.framewin, COLOR_PAIR(0));
    wborder(ttyclock.framewin, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ');
    werase(ttyclock.framewin);
    wrefresh(ttyclock.framewin);

    if(ttyclock.option.date){
        wbkgdset(ttyclock.datewin, COLOR_PAIR(0));
        wborder(ttyclock.datewin, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ');
        werase(ttyclock.datewin);
        wrefresh(ttyclock.datewin);
    }

    /* Frame win move */
    mvwin(ttyclock.framewin, (ttyclock.geo.x = x), (ttyclock.geo.y = y));
    wresize(ttyclock.framewin, (ttyclock.geo.h = h), (ttyclock.geo.w = w));

    /* Date win move */
    if (ttyclock.option.date){
        mvwin(ttyclock.datewin,
              ttyclock.geo.x + ttyclock.geo.h - 1,
              ttyclock.geo.y + (ttyclock.geo.w / 2) - (strlen(ttyclock.date.datestr) / 2) - 1);
        wresize(ttyclock.datewin, DATEWINH, strlen(ttyclock.date.datestr) + 2);

        if (ttyclock.option.box) {
            box(ttyclock.datewin, 0, 0);
        }
    }

    if (ttyclock.option.box) {
        box(ttyclock.framewin, 0, 0);
    }

    wrefresh(ttyclock.framewin);
    wrefresh(ttyclock.datewin);
    return;
}

/* Useless buf fun :) */
void clock_rebound(void)
{
    if(!ttyclock.option.rebound)
        return;

    if(ttyclock.geo.x < 1)
        ttyclock.geo.a = 1;
    if(ttyclock.geo.x > (LINES - ttyclock.geo.h - DATEWINH))
        ttyclock.geo.a = -1;
    if(ttyclock.geo.y < 1)
        ttyclock.geo.b = 1;
    if(ttyclock.geo.y > (COLS - ttyclock.geo.w - 1))
        ttyclock.geo.b = -1;

    clock_move(ttyclock.geo.x + ttyclock.geo.a,
               ttyclock.geo.y + ttyclock.geo.b,
               ttyclock.geo.w,
               ttyclock.geo.h);

    return;
}

void update_hour(void)
{
    int ihour;
    char tmpstr[128];

    ttyclock.lt = time(NULL);
    ttyclock.tm = localtime(&(ttyclock.lt));
    if(ttyclock.option.utc) {
        ttyclock.tm = gmtime(&(ttyclock.lt));
    }

    ihour = ttyclock.tm->tm_hour;

    if(ttyclock.option.twelve)
        ttyclock.meridiem = ((ihour >= 12) ? PMSIGN : AMSIGN);
    else
        ttyclock.meridiem = "\0";

    /* Manage hour for twelve mode */
    ihour = ((ttyclock.option.twelve && ihour > 12) ? (ihour - 12) : ihour);
    ihour = ((ttyclock.option.twelve && !ihour) ? 12 : ihour);

    /* Set hour */
    ttyclock.date.hour[0] = ihour / 10;
    ttyclock.date.hour[1] = ihour % 10;

    /* Set minutes */
    ttyclock.date.minute[0] = ttyclock.tm->tm_min / 10;
    ttyclock.date.minute[1] = ttyclock.tm->tm_min % 10;

    /* Set date string */
    strcpy(ttyclock.date.old_datestr, ttyclock.date.datestr);
    strftime(tmpstr,
             sizeof(tmpstr),
             ttyclock.option.format,
             ttyclock.tm);
    sprintf(ttyclock.date.datestr, "%s%s", tmpstr, ttyclock.meridiem);

    /* Set seconds */
    ttyclock.date.second[0] = ttyclock.tm->tm_sec / 10;
    ttyclock.date.second[1] = ttyclock.tm->tm_sec % 10;

    return;
}

void draw_number(int n, int x, int y)
{
    int i, sy = y;

    for(i = 0; i < 30; ++i, ++sy){
        if(sy == y + 6){
            sy = y;
            ++x;
        }

        if(ttyclock.option.bold)
            wattron(ttyclock.framewin, A_BLINK);
        else
            wattroff(ttyclock.framewin, A_BLINK);

        wbkgdset(ttyclock.framewin, COLOR_PAIR(number[n][i/2]));
        mvwaddch(ttyclock.framewin, x, sy, ' ');
    }
    wrefresh(ttyclock.framewin);

    return;
}

void draw_clock(void)
{
    if(ttyclock.option.date && !ttyclock.option.rebound &&
        strcmp(ttyclock.date.datestr, ttyclock.date.old_datestr) != 0)
    {
        clock_move(ttyclock.geo.x,
                   ttyclock.geo.y,
                   ttyclock.geo.w,
                   ttyclock.geo.h);
    }

    /* Draw hour numbers */
    draw_number(ttyclock.date.hour[0], 1, 1);
    draw_number(ttyclock.date.hour[1], 1, 8);
    chtype dotcolor = COLOR_PAIR(1);
    if(ttyclock.option.blink && time(NULL) % 2 == 0)
        dotcolor = COLOR_PAIR(2);

    /* 2 dot for number separation */
    wbkgdset(ttyclock.framewin, dotcolor);
    mvwaddstr(ttyclock.framewin, 2, 16, "  ");
    mvwaddstr(ttyclock.framewin, 4, 16, "  ");

    /* Draw minute numbers */
    draw_number(ttyclock.date.minute[0], 1, 20);
    draw_number(ttyclock.date.minute[1], 1, 27);

    /* Draw the date */
    if(ttyclock.option.bold)
        wattron(ttyclock.datewin, A_BOLD);
    else
        wattroff(ttyclock.datewin, A_BOLD);

    if(ttyclock.option.date){
        wbkgdset(ttyclock.datewin, (COLOR_PAIR(2)));
        mvwprintw(ttyclock.datewin, (DATEWINH / 2), 1, "%s", ttyclock.date.datestr);
        wrefresh(ttyclock.datewin);
    }

    /* Draw second if the option is enabled */
    if(ttyclock.option.second){
        /* Again 2 dor for number separation */
        wbkgdset(ttyclock.framewin, dotcolor);
        mvwaddstr(ttyclock.framewin, 2, NORMFRAMEW, "  ");
        mvwaddstr(ttyclock.framewin, 4, NORMFRAMEW, "  ");

        /* Draw second numbers */
        draw_number(ttyclock.date.second[0], 1, 39);
        draw_number(ttyclock.date.second[1], 1, 46);
    }

    return;
}

void set_second(void)
{
    int new_w = (((ttyclock.option.second = !ttyclock.option.second)) ? SECFRAMEW : NORMFRAMEW);
    int y_adj;

    for(y_adj = 0; (ttyclock.geo.y - y_adj) > (COLS - new_w - 1); ++y_adj);

    clock_move(ttyclock.geo.x, (ttyclock.geo.y - y_adj), new_w, ttyclock.geo.h);

    set_center(ttyclock.option.center);

    return;
}

void set_center(bool b)
{
    if((ttyclock.option.center = b))
    {
        ttyclock.option.rebound = false;

        clock_move((LINES / 2 - (ttyclock.geo.h / 2)),
                   (COLS  / 2 - (ttyclock.geo.w / 2)),
                   ttyclock.geo.w,
                   ttyclock.geo.h);
    }

    return;
}

void set_box(bool b)
{
    ttyclock.option.box = b;

    wbkgdset(ttyclock.framewin, COLOR_PAIR(0));
    wbkgdset(ttyclock.datewin, COLOR_PAIR(0));

    if(ttyclock.option.box) {
        wbkgdset(ttyclock.framewin, COLOR_PAIR(0));
        wbkgdset(ttyclock.datewin, COLOR_PAIR(0));
        box(ttyclock.framewin, 0, 0);
        box(ttyclock.datewin, 0, 0);
    } else {
        wborder(ttyclock.framewin, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ');
        wborder(ttyclock.datewin, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ');
    }

    wrefresh(ttyclock.datewin);
    wrefresh(ttyclock.framewin);
}

void key_event(void)
{
    int i, c;

    struct timespec length = { ttyclock.option.delay, ttyclock.option.nsdelay };

    fd_set rfds;
    FD_ZERO(&rfds);
    FD_SET(STDIN_FILENO, &rfds);

    if(ttyclock.option.screensaver) {
        c = wgetch(stdscr);
        if(c != ERR && ttyclock.option.noquit == false) {
            ttyclock.running = false;
        }
        else {
            nanosleep(&length, NULL);
            for(i = 0; i < 8; ++i) {
                if(c == (i + '0')) {
                    ttyclock.option.color = i;
                    init_pair(1, ttyclock.bg, i);
                    init_pair(2, i, ttyclock.bg);
                }
            }
        }
        return;
    }

    switch (c = wgetch(stdscr)) {
        case KEY_RESIZE:
            endwin();
            init();
            break;

        case KEY_UP:
        case 'k':
        case 'K':
            if(ttyclock.geo.x >= 1 && !ttyclock.option.center)
                clock_move(ttyclock.geo.x - 1, ttyclock.geo.y, ttyclock.geo.w, ttyclock.geo.h);
            break;

        case KEY_DOWN:
        case 'j':
        case 'J':
            if(ttyclock.geo.x <= (LINES - ttyclock.geo.h - DATEWINH) && !ttyclock.option.center)
                clock_move(ttyclock.geo.x + 1, ttyclock.geo.y, ttyclock.geo.w, ttyclock.geo.h);
            break;

        case KEY_LEFT:
        case 'h':
        case 'H':
            if(ttyclock.geo.y >= 1 && !ttyclock.option.center)
                clock_move(ttyclock.geo.x, ttyclock.geo.y - 1, ttyclock.geo.w, ttyclock.geo.h);
            break;

        case KEY_RIGHT:
        case 'l':
        case 'L':
            if(ttyclock.geo.y <= (COLS - ttyclock.geo.w - 1) && !ttyclock.option.center)
                clock_move(ttyclock.geo.x, ttyclock.geo.y + 1, ttyclock.geo.w, ttyclock.geo.h);
            break;

        case 'q':
        case 'Q':
            if(ttyclock.option.noquit == false)
                ttyclock.running = false;
            break;

        case 's':
        case 'S':
            set_second();
            break;

        case 't':
        case 'T':
            ttyclock.option.twelve = !ttyclock.option.twelve;
            /* Set the new ttyclock.date.datestr to resize date window */
            update_hour();
            clock_move(ttyclock.geo.x, ttyclock.geo.y, ttyclock.geo.w, ttyclock.geo.h);
            break;

        case 'c':
        case 'C':
            set_center(!ttyclock.option.center);
            break;

        case 'b':
        case 'B':
            ttyclock.option.bold = !ttyclock.option.bold;
            break;

        case 'r':
        case 'R':
            ttyclock.option.rebound = !ttyclock.option.rebound;
            if(ttyclock.option.rebound && ttyclock.option.center)
                ttyclock.option.center = false;
            break;

        case 'x':
        case 'X':
            set_box(!ttyclock.option.box);
            break;

            case '0': case '1': case '2': case '3':
            case '4': case '5': case '6': case '7':
                i = c - '0';
                ttyclock.option.color = i;
                init_pair(1, ttyclock.bg, i);
                init_pair(2, i, ttyclock.bg);
                break;

        default:
            pselect(1, &rfds, NULL, NULL, &length, NULL);
    }

    return;
}

int main(int argc, char **argv)
{
    //int c;

    // alloc ttyclock
    memset(&ttyclock, 0x00, sizeof(ttyclock_t));

    ttyclock.option.date = true;

    // default date
    strncpy(ttyclock.option.format, "%F", sizeof(ttyclock.option.format));
    // default color
    ttyclock.option.color = 3;
    // default delay
    ttyclock.option.delay = 1;
    ttyclock.option.nsdelay = 0;
    ttyclock.option.blink = false;

    atexit(cleanup);

    /*
     *　on hold
    while ((c = getopt(argc, argv, "iuvsScbtrhBxnDC:f:d:T:a:")) != -1)
    {
        ;
    }
     */

    init();
    attron(A_BLINK);
    while(ttyclock.running) {
        clock_rebound();
        update_hour();
        draw_clock();
        key_event();
    }

    endwin();

    return 0;
}
