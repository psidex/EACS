#include <stdio.h>
#define TRAY_WINAPI 1
#include "tray.h"
#include "configs.h"

struct tray tray_app;
int config_count;
struct e_apo_config *config_array;

void config_clicked(struct tray_menu *item) {
    struct e_apo_config *current_config = item->context;

    // Update the tray and the struct
    item->checked = !item->checked;
    current_config->checked = !current_config->checked;

    config_write_to_file(config_count, config_array);
    tray_update(&tray_app);
}

void quit_app(struct tray_menu *item) {
    tray_exit();
}

void tray_menu_populate(struct tray_menu *tray_menu_items) {
    // Takes a pointer to a tray_menu struct and inserts all needed items
    tray_menu_items[0] = (struct tray_menu) {"E-APO-Config-Switcher", 1, 0, NULL, NULL};
    tray_menu_items[1] = (struct tray_menu) {"-", 0, 0, NULL, NULL};
    tray_menu_items[config_count+2] = (struct tray_menu) {"-", 0, 0, NULL, NULL};
    tray_menu_items[config_count+3] = (struct tray_menu) {"Quit", 0, 0, quit_app, NULL};
    tray_menu_items[config_count+4] = (struct tray_menu) {NULL, 0, 0, NULL, NULL};

    for(int i = 0; i < config_count; i++) {
        // i + 2 as the first 2 indexes are already used
        tray_menu_items[i + 2] = (struct tray_menu) {
                config_array[i].name,
                0,
                0,
                config_clicked,
                // Pass a pointer to the e_apo_config struct as context
                &config_array[i]
        };
    }
}

// See CMakeLists.txt for WinMain usage reason
int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR szCmdLine, int iCmdShow ) {
    config_count = config_file_count();
    config_array = malloc(config_count * sizeof(struct e_apo_config));
    struct tray_menu *tray_menu_items = malloc((config_count + 5) * sizeof(struct tray_menu));

    config_populate_array(config_array);
    tray_menu_populate(tray_menu_items);

    tray_app.icon = "icon.ico";
    tray_app.menu = tray_menu_items;

    if (tray_init(&tray_app) < 0) {return 1;}
    while (tray_loop(1) == 0) {}

    free(config_array);
    free(tray_menu_items);

    return 0;
}
