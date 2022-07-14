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

}
