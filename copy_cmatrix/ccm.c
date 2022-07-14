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

typedef struct cmatrix {
    int val;
    bool is_head;
} cmatrix;

// global variables
cmatrix **matrix = (cmatrix **) NULL;
int console = 0;

int va_system(char *str, ...) {
    va_list ap;
    char buf[133];

    va_start(ap, str);
    vsnprintf(buf, sizeof(buf), str, ap);
    va_end(ap);
    return system(buf);
}

void c_die(char *msg, ...) {
    
    va_list ap;

    curs_set(1);
    clear();
    refresh();
    resetty();
    endwin();

    if (console) {
#ifdef HAVE_CONSOLECHARS
        va_system("consolechars -d");
    }


}

void nmalloc(size_t howmuch) {
    void *r;

    if (!(r = malloc(howmuch))) {
        c_die("CMatrix: malloc: out of memory!");
    }

    return r;
}

void var_init() {
    int i, j;

    if (matrix != NULL) {
        free(matrix[0]);
        free(matrix);
    }

    matrix = nmalloc(sizeof())
}

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
    iint mcolor = COLOR_GREEN;
    int rainbow = 0;
    int lambda = 0;
    int pause = 0;
    int classic = 0;
    int changes = 0;
    char *msg = "";
    char *tty = NULL;
    */

    int highnum = 0;
    int randmin = 0;
    int randnum = 0;

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
    
    // set up values for random number generation
    randmin  = 33;
    highnum = 123;
    randnum = highnum - randmin; 

    //var_init();
}
