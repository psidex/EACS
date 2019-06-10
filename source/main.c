#include <stdio.h>
#define TRAY_WINAPI 1
#include "tray.h"
#include "helpers.h"

struct tray tray_app;
int config_file_count;
char **config_file_names;

static const char config_file_base_path[] = "E-APO-Config-Switcher\\E-APO-Config-Files\\";

void config_file_clicked(struct tray_menu *item) {
    // When a config file in the tray menu is clicked
    item->checked = !item->checked;

    // TODO: Sort out lengths
    char include_text[999] = "Include: ";
    strcat_s(include_text, 999, config_file_base_path);
    strcat_s(include_text, FILE_NAME_LEN_LIMIT, item->text);
    printf("include_text: %s\n", include_text);

    tray_update(&tray_app);
}

void quit_app(struct tray_menu *item) {
    tray_exit();
}

void populate_tray_menu_items(struct tray_menu *tray_menu_items) {
    struct tray_menu spacer_item;
    spacer_item.text = "-";

    tray_menu_items[0] = (struct tray_menu) {"E-APO-Config-Switcher", 1, 0, NULL, NULL};
    tray_menu_items[1] = spacer_item;

    for(int i = 0; i < config_file_count; i++) {
        tray_menu_items[2+i] = (struct tray_menu) {config_file_names[i], 0, 0, config_file_clicked};
    }

    tray_menu_items[config_file_count+2] = spacer_item;
    tray_menu_items[config_file_count+3] = (struct tray_menu) {"Quit", 0, 0, quit_app, NULL};
    tray_menu_items[config_file_count+4] = (struct tray_menu) {NULL, 0, 0, NULL, NULL};
}

int main() {
    // Setup and get the array of config file names
    config_file_count = get_config_file_count();
    config_file_names = malloc(config_file_count * sizeof(char*));
    populate_config_file_names(config_file_names);

    for(int i = 0; i < config_file_count; i++) {
        // TODO: Some sort of logging
        printf("Config file loaded: %s\n", config_file_names[i]);
    }

    // Setup tray menu array
    // +5 because there are 5 menus items other than the config file names
    struct tray_menu *tray_menu_items = malloc((config_file_count+5) * sizeof(struct tray_menu));
    populate_tray_menu_items(tray_menu_items);

    tray_app.icon = "icon.ico";
    tray_app.menu = tray_menu_items;

    // Init and start tray app
    if (tray_init(&tray_app) < 0) {return 1;}
    while (tray_loop(1) == 0) {}

    // Free up allocated memory
    for(int i = 0; i < config_file_count; i++) free(config_file_names[i]);
    free(config_file_names);

    return 0;
}
