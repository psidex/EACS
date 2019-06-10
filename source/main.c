#include <stdio.h>
#define TRAY_WINAPI 1
#include "tray.h"
#include "helpers.h"

struct tray tray_app;
int config_file_count;
struct e_apo_config *e_apo_configs;

// The file that contains includes to config files
const char master_config_path[] = "../config.txt";

void write_current_config() {
    // For each checked config, write the include text to the master config file
    // TODO: Switch to fopen_s?
    FILE *fp = fopen(master_config_path, "w");
    for(int i = 0; i < config_file_count; i++) {
        if (e_apo_configs[i].checked) {
            fputs(e_apo_configs[i].include_text, fp);
            fputs("\n", fp);  // E-APO doesn't like CR LF, only LF
        }
    }
    fclose(fp);
}

void config_clicked(struct tray_menu *item) {
    struct e_apo_config *current_config = item->context;

    // Update the tray and the struct
    item->checked = !item->checked;
    current_config->checked = !current_config->checked;

    write_current_config();
    tray_update(&tray_app);
}

void quit_app(struct tray_menu *item) {
    tray_exit();
}

void populate_tray_menu(struct tray_menu *tray_menu_items) {
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
                // Pass a pointer to the e_apo_config struct as context
                &e_apo_configs[i]
        };
    }
}

// See CMakeLists.txt for WinMain usage reason
int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR szCmdLine, int iCmdShow ) {
    config_file_count = get_config_file_count();

    e_apo_configs = malloc(config_file_count * sizeof(struct e_apo_config));
    struct tray_menu *tray_menu_items = malloc((config_file_count + 5) * sizeof(struct tray_menu));

    populate_e_apo_configs(e_apo_configs);
    populate_tray_menu(tray_menu_items);

    tray_app.icon = "icon.ico";
    tray_app.menu = tray_menu_items;

    // Init and start tray app
    if (tray_init(&tray_app) < 0) {return 1;}
    while (tray_loop(1) == 0) {}

    // Free up allocated memory
    free(e_apo_configs);
    free(tray_menu_items);

    return 0;
}
