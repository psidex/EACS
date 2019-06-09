#include <stdio.h>
#define TRAY_WINAPI 1
#include "tray/tray.h"

struct tray tray;

void toggle_menu_item(struct tray_menu *item) {
    item->checked = !item->checked;
    printf("%s was clicked", item->text);
    tray_update(&tray);
}

void quit_app(struct tray_menu *item) {
    tray_exit();
}

struct tray tray = {
        .icon = "icon.ico",
        .menu = (struct tray_menu[]) {
            // Text, Disabled, Checked, Callback, Optional "context pointer"
            {"E-APO-Config-Switcher", 1, 0, NULL, NULL},
            {"-", 0, 0, NULL, NULL},
            {"Toggle me", 0, 0, toggle_menu_item, NULL},
            {"-", 0, 0, NULL, NULL},
            {"Quit", 0, 0, quit_app, NULL},
            {NULL, 0, 0, NULL, NULL}
            },
};

int main() {
    if (tray_init(&tray) < 0) {return 1;}
    // tray_loop(blocking_true_or_false)
    while (tray_loop(1) == 0) {}
    return 0;
}
