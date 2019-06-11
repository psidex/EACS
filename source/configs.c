#include <stdio.h>
#include <windows.h>
#include "configs.h"

// Glob to get all config files
const char config_files_glob[] = "config-files\\*.txt";

// The master config file that Equalizer APO reads
const char config_master_path[] = "../config.txt";

// config_files_base is relative to the master config file
const char config_files_base[] = "E-APO-Config-Switcher\\config-files\\";

void config_write_to_file(int config_count, struct e_apo_config *e_apo_configs) {
    // For each checked config, write the include_text to the master config file
    // TODO: Switch to fopen_s?, handle err
    FILE *fp = fopen(config_master_path, "w");
    for(int i = 0; i < config_count; i++) {
        if (e_apo_configs[i].checked) fputs(e_apo_configs[i].include_text, fp);
    }
    fclose(fp);
}

int config_file_count() {
    // https://stackoverflow.com/a/3176223/6396652
    int file_count = 0;
    WIN32_FIND_DATA data;
    HANDLE hFind = FindFirstFile(config_files_glob, &data);
    if (hFind != INVALID_HANDLE_VALUE) {
        do {file_count++;} while (FindNextFile(hFind, &data));
        FindClose(hFind);
    }
    return file_count;
}

void config_populate_array(struct e_apo_config *e_apo_configs) {
    // Takes a pointer to an array of e_apo_config and inserts the file info
    int current_file_number = 0;

    WIN32_FIND_DATA data;
    HANDLE hFind = FindFirstFile(config_files_glob, &data);
    if (hFind != INVALID_HANDLE_VALUE) {
        do {
            char filename_no_ext[MAX_FILE_NAME];
            strcpy_s(filename_no_ext, MAX_FILE_NAME, data.cFileName);
            // https://stackoverflow.com/a/1726318/6396652
            // Set the 4th to last char to the end of the string, removing the ".txt"
            // The glob used for FindFirstFile assures this wont break things
            filename_no_ext[strlen(filename_no_ext)-4] = 0;

            char include_text[MAX_INCLUDE_TEXT];
            // E-APO doesn't like CR LF, only LF
            sprintf_s(include_text, MAX_INCLUDE_TEXT, "Include: %s%s\n", config_files_base, data.cFileName);

            strcpy_s(e_apo_configs[current_file_number].file_name, MAX_FILE_NAME, filename_no_ext);
            strcpy_s(e_apo_configs[current_file_number].include_text, MAX_INCLUDE_TEXT, include_text);
            e_apo_configs[current_file_number].checked = 0;

            current_file_number++;
        } while (FindNextFile(hFind, &data));
        FindClose(hFind);
    }
}
