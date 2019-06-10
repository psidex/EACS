#include <windows.h>
#include "helpers.h"

static const char config_files_glob[] = "E-APO-Config-Files\\*.txt";

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

void populate_config_file_names(char **config_file_names) {
    int current_file_number = 0;
    HANDLE hFind;
    WIN32_FIND_DATA data;

    hFind = FindFirstFile(config_files_glob, &data);
    if (hFind != INVALID_HANDLE_VALUE) {
        do {
            config_file_names[current_file_number] = malloc(FILE_NAME_LEN_LIMIT * sizeof(char));
            strcpy_s(config_file_names[current_file_number], FILE_NAME_LEN_LIMIT, data.cFileName);
            current_file_number++;
        } while (FindNextFile(hFind, &data));
        FindClose(hFind);
    }
}
