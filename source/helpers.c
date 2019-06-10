#include <windows.h>
#include "helpers.h"

const char config_files_glob[] = "E-APO-Config-Files\\*.txt";
// config_files_base is relative to the master config.txt
const char config_files_base[] = "E-APO-Config-Switcher\\E-APO-Config-Files\\";

int get_config_file_count() {
    int file_count = 0;
    HANDLE hFind;
    WIN32_FIND_DATA data;

    // https://stackoverflow.com/a/3176223/6396652
    hFind = FindFirstFile(config_files_glob, &data);
    if (hFind != INVALID_HANDLE_VALUE) {
        do {
            file_count++;
        } while (FindNextFile(hFind, &data));
        FindClose(hFind);
    }

    return file_count;
}

void populate_e_apo_configs(struct e_apo_config *e_apo_configs) {
    int current_file_number = 0;
    HANDLE hFind;
    WIN32_FIND_DATA data;

    hFind = FindFirstFile(config_files_glob, &data);
    if (hFind != INVALID_HANDLE_VALUE) {
        do {
            char include_text[MAX_INCLUDE_TEXT] = "Include: ";
            strcat_s(include_text, MAX_FILE_NAME, config_files_base);
            strcat_s(include_text, MAX_FILE_NAME, data.cFileName);

            strcpy_s(e_apo_configs[current_file_number].file_name, MAX_FILE_NAME, data.cFileName);
            strcpy_s(e_apo_configs[current_file_number].include_text, MAX_INCLUDE_TEXT, include_text);
            e_apo_configs[current_file_number].checked = 0;

            current_file_number++;
        } while (FindNextFile(hFind, &data));
        FindClose(hFind);
    }
}
