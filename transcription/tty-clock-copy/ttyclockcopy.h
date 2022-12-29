//
// Created by stanaka on 22/12/24.
//

#ifndef TTY_CLOCK_COPY_TTYCLOCKCOPY_H
#define TTY_CLOCK_COPY_TTYCLOCKCOPY_H

#include <stdbool.h>
#include <ncurses.h>
#include <time.h>
#include <stdlib.h>
#include <string.h>
#include <locale.h>
#include <errno.h>
#include <assert.h>
#include <signal.h>
#include <unistd.h>

// Macro
#define NORMFRAMEW  35
#define SECFRAMEW   54
#define AMSIGN      " [AM]"
#define PMSIGN      " [PM]"
#define DATEWINH    3

typedef struct
{
    bool running;

    // terminal variables
    SCREEN *ttyscr;
    char *tty;
    int bg;

    // Running option
    struct
    {
        bool second;
        bool screensaver;
        bool twelve;
        bool center;
        bool rebound;
        bool date;
        bool utc;
        bool box;
        bool noquit;
        char format[100];
        int color;
        bool bold;
        long delay;
        bool blink;
        long nsdelay;
    } option;

    // clock geometry
    struct
    {
        int x, y, w, h;
        int a, b;
    } geo;

    // date content
    struct
    {
        unsigned int hour[2];
        unsigned int minute[2];
        unsigned int second[2];
        char datestr[256];
        char old_datestr[256];
    } date;

    // time.h utils
    struct tm *tm;
    time_t lt;

    // clock member
    char *meridiem;
    WINDOW *framewin;
    WINDOW *datewin;

} ttyclock_t;

/* Prototypes */
void init(void);
void signal_handler(int signal);
void update_hour(void);
void draw_number(int n, int x, int y);
void draw_clock(void);
void clock_move(int x, int y, int w, int h);
void set_second(void);
void set_center(bool b);
void set_box(bool b);
void key_event(void);

// global variable
ttyclock_t ttyclock;

/* Number matrix */
const bool number[][15] =
        {
                {1,1,1,1,0,1,1,0,1,1,0,1,1,1,1}, /* 0 */
                {0,0,1,0,0,1,0,0,1,0,0,1,0,0,1}, /* 1 */
                {1,1,1,0,0,1,1,1,1,1,0,0,1,1,1}, /* 2 */
                {1,1,1,0,0,1,1,1,1,0,0,1,1,1,1}, /* 3 */
                {1,0,1,1,0,1,1,1,1,0,0,1,0,0,1}, /* 4 */
                {1,1,1,1,0,0,1,1,1,0,0,1,1,1,1}, /* 5 */
                {1,1,1,1,0,0,1,1,1,1,0,1,1,1,1}, /* 6 */
                {1,1,1,0,0,1,0,0,1,0,0,1,0,0,1}, /* 7 */
                {1,1,1,1,0,1,1,1,1,1,0,1,1,1,1}, /* 8 */
                {1,1,1,1,0,1,1,1,1,0,0,1,1,1,1}, /* 9 */
        };

#endif //TTY_CLOCK_COPY_TTYCLOCKCOPY_H
