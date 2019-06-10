#include <stdio.h>
#define TRAY_WINAPI 1
#include "tray.h"
#include "helpers.h"

struct tray tray_app;

void config_clicked(struct tray_menu *item) {
    // When a config file in the tray menu is clicked
    item->checked = !item->checked;

    struct e_apo_config *current_config = item->context;
    printf("include_text: %s\n", current_config->include_text);

    tray_update(&tray_app);
}

void quit_app(struct tray_menu *item) {
    tray_exit();
}

int main() {
    int config_file_count = get_config_file_count();

    struct e_apo_config *e_apo_configs = malloc(config_file_count * sizeof(struct e_apo_config));
    populate_e_apo_configs(e_apo_configs);

    struct tray_menu *tray_menu_items = malloc((config_file_count + 5) * sizeof(struct tray_menu));
    tray_menu_items[0] = (struct tray_menu) {"E-APO-Config-Switcher", 1, 0, NULL, NULL};
    tray_menu_items[1] = (struct tray_menu) {"-", 0, 0, NULL, NULL};
    tray_menu_items[config_file_count+2] = (struct tray_menu) {"-", 0, 0, NULL, NULL};
    tray_menu_items[config_file_count+3] = (struct tray_menu) {"Quit", 0, 0, quit_app, NULL};
    tray_menu_items[config_file_count+4] = (struct tray_menu) {NULL, 0, 0, NULL, NULL};

    for(int i = 0; i < config_file_count; i++) {
        // i + 2 as the first 2 indexes are already used
        tray_menu_items[i + 2] = (struct tray_menu) {
            e_apo_configs[i].file_name,
            0,
            0,
            config_clicked,
            // Pass the e_apo_config struct as context
            &e_apo_configs[i]
        };
    }

    tray_app.icon = "icon.ico";
    tray_app.menu = tray_menu_items;

    // TODO: Logging to file
    for(int i = 0; i < config_file_count; i++) {
        printf("Config file loaded: %s\n", e_apo_configs[i].file_name);
    }

    // Init and start tray app
    if (tray_init(&tray_app) < 0) {return 1;}
    while (tray_loop(1) == 0) {}

    // Free up allocated memory
    free(e_apo_configs);
    free(tray_menu_items);

    return 0;
}
