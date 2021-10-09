#include <stdio.h>

typedef struct {
    double x;
    double y;
} Point;

void draw_line(Point start_p, Point end_p)
{
    fprintf(stderr, "draw line: (%f, %f)-(%f,%f)\n", start_p.x, start_p.y, end_p.x, end_p.y);
}

void draw_polyline(int npoints, Point *point)
{
    for (int i=0; i<npoints; i++) {
        fprintf(stderr, "[i]..(%f, %f)\n", point[i].x, point[i].y);
    }
}

int main(void) 
{
    draw_line((Point){.x = 10, .y = 10}, (Point){.x = 10, .y = 20});

    draw_polyline(5, (Point[]){
                (Point){.x = 1, .y = 1},
                (Point){.x = 2, .y = 2},
                (Point){.x = 3, .y = 3},
                (Point){.x = 4, .y = 4},
                (Point){.x = 5, .y = 5},
            });
}
