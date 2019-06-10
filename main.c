#include <stdio.h>
#include "helpers.h"
#define TRAY_WINAPI 1
#include "tray/tray.h"

struct tray tray_app;

void toggle_menu_item(struct tray_menu *item) {
    item->checked = !item->checked;
    printf("%s was clicked\n", item->text);
    tray_update(&tray_app);
}

void quit_app(struct tray_menu *item) {
    tray_exit();
}

void populate_tray_menu_items(struct tray_menu *tray_menu_items, char **config_file_names, int file_count) {
    struct tray_menu spacer_item;
    spacer_item.text = "-";

    tray_menu_items[0] = (struct tray_menu) {"E-APO-Config-Switcher", 1, 0, NULL, NULL};
    tray_menu_items[1] = spacer_item;

    for(int i = 0; i < file_count; i++) {
        tray_menu_items[2+i] = (struct tray_menu) {config_file_names[i], 0, 0, toggle_menu_item};
    }

    tray_menu_items[file_count+2] = spacer_item;
    tray_menu_items[file_count+3] = (struct tray_menu) {"Quit", 0, 0, quit_app, NULL};
    tray_menu_items[file_count+4] = (struct tray_menu) {NULL, 0, 0, NULL, NULL};
}

int main() {
    // Get an array of config file names
    int file_count = get_config_file_count();
    char **config_file_names = malloc(file_count * sizeof(char*));
    populate_config_file_names(config_file_names);

    for(int i = 0; i < file_count; i++) {
        // TODO: Some sort of logging
        printf("Config file loaded: %s\n", config_file_names[i]);
    }

    // Setup tray menu array
    struct tray_menu *tray_menu_items = malloc((file_count+5) * sizeof(struct tray_menu));
    populate_tray_menu_items(tray_menu_items, config_file_names, file_count);

    tray_app.icon = "icon.ico";
    tray_app.menu = tray_menu_items;

    // Init and start tray app
    if (tray_init(&tray_app) < 0) {return 1;}
    while (tray_loop(1) == 0) {}

    // Free up allocated memory
    for(int i = 0; i < file_count; i++) free(config_file_names[i]);
    free(config_file_names);

    return 0;
}
