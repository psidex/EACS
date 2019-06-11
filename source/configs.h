#ifndef E_APO_CONFIG_SWITCHER_CONFIGS_H
#define E_APO_CONFIG_SWITCHER_CONFIGS_H

#define MAX_FILE_NAME 255
// Max path len + 9 for "Include: " and + 1 for '\n'
#define MAX_INCLUDE_TEXT 260 + 10

struct e_apo_config {
    char file_name[MAX_FILE_NAME];
    char include_text[MAX_INCLUDE_TEXT];
    int checked;
};

void config_write_to_file(int config_count, struct e_apo_config *e_apo_configs);
int config_file_count();
void config_populate_array(struct e_apo_config *e_apo_configs);

#endif //E_APO_CONFIG_SWITCHER_CONFIGS_H
