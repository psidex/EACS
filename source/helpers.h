#ifndef E_APO_CONFIG_SWITCHER_HELPERS_H
#define E_APO_CONFIG_SWITCHER_HELPERS_H

#define MAX_FILE_NAME 255
// Max path len + 9 for "Include: "
#define MAX_INCLUDE_TEXT 260 + 9

struct e_apo_config {
    char file_name[MAX_FILE_NAME];
    char include_text[MAX_INCLUDE_TEXT];
    int checked;
};

int get_config_file_count();
void populate_e_apo_configs(struct e_apo_config *e_apo_configs);

#endif //E_APO_CONFIG_SWITCHER_HELPERS_H
