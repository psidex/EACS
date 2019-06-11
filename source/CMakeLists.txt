cmake_minimum_required(VERSION 3.13)
project(E_APO_Config_Switcher C)
set(CMAKE_C_STANDARD 99)

include_directories(../tray)

# WIN32 = Don't show a console window when running
# This requires using WinMain instead of main - https://stackoverflow.com/a/13871657/6396652
add_executable(E_APO_Config_Switcher WIN32 main.c configs.c configs.h)

# Move icon over to binary dir so it is seen by the bin
# configure_file means it will copy if it changes
configure_file(icon.ico ${CMAKE_CURRENT_BINARY_DIR} COPYONLY)

# Move config files into binary dir
# https://stackoverflow.com/a/24311917/6396652
# TODO: Find method that copies when the dir changes
file(GLOB E_APO_CONFIG_FILES "config-files/*.txt")
file(COPY ${E_APO_CONFIG_FILES} DESTINATION ${CMAKE_CURRENT_BINARY_DIR}/config-files)
