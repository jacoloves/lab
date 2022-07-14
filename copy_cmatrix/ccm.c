#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <string.h>
#include <time.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <signal.h>
#include <locale.h>
#include <curses.h>

#ifdef HAVE_UNISTD_H
#include <unistd.h>
#endif

int main(int argc, char *argv[]) {
    /*
    int i, y, z, opychr, keyperss;
    int j = 0;
    int count = 0;
    int screensaver = 0;
    int asynch = 0;
    int bold = 0;
    int force = 0;
    int firstcoldone = 0;
    int random = 0;
    int update = 4;
    int highnum = 0;
    iint mcolor = COLOR_GREEN;
    int rainbow = 0;
    int lambda = 0;
    int randnum = 0;
    int randmin = 0;
    int pause = 0;
    int classic = 0;
    int changes = 0;
    char *msg = "";
    char *tty = NULL;
    */

    srand((unsigned) time(NULL));
    setlocale(LC_ALL, "");

    // process start
    initscr();
    savetty();
    nonl();
    cbreak();
    noecho();
    timeout(0);
    leaveok(stdscr, TRUE);
    curs_set(0);
    
    // check color
    if (has_colors()) {
        start_color();
#ifdef HAVE_USE_DEFAULT_COLORS
        if (use_default_color() != ERR) {
            init_pair(COLOR_BLACK, -1, -1);
            init_pair(COLOR_GREEN, COLOR_GREEN, -1);
            init_pair(COLOR_WHITE, COLOR_WHITE, -1);
            init_pair(COLOR_RED, COLOR_RED, -1);
            init_pair(COLOR_CYAN, COLOR_CYAN, -1);
            init_pair(COLOR_MAGENTA, COLOR_MAGENTA, -1);
            init_pair(COLOR_BLUE, COLOR_BLUE, -1);
            init_pair(COLOR_YELLOW, COLOR_YELLOW, -1);
        } else {
#else
        {
#endif
            init_pair(COLOR_BLACK, COLOR_BLACK, COLOR_BLACK);
            init_pair(COLOR_GREEN, COLOR_GREEN, COLOR_BLACK);
            init_pair(COLOR_WHITE, COLOR_WHITE, COLOR_BLACK);
            init_pair(COLOR_RED, COLOR_RED, COLOR_RED);
            init_pair(COLOR_CYAN, COLOR_CYAN, COLOR_CYAN);
            init_pair(COLOR_MAGENTA, COLOR_MAGENTA, COLOR_MAGENTA);
            init_pair(COLOR_BLUE, COLOR_BLUE, COLOR_BLUE);
            init_pair(COLOR_YELLOW, COLOR_YELLOW, COLOR_YELLOW);
        }
    }
}
