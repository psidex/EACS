project(E_APO_Config_Switcher C)
cmake_minimum_required(VERSION 3.13)
set(CMAKE_C_STANDARD 99)

include_directories(../tray)
set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -Wall")

if(CMAKE_BUILD_TYPE MATCHES Debug)
    # Doen't add the WIN32 option so stdout will show in a console
    add_executable(E_APO_Config_Switcher main.c configs.c configs.h)
else()
    # WIN32 = Don't show a console window when running
    # This requires using WinMain instead of main - https://stackoverflow.com/a/13871657/6396652
    add_executable(E_APO_Config_Switcher WIN32 main.c configs.c configs.h)
endif()

# Move files over to binary dir so they can be seen by the bin
file(COPY icon.ico DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
file(GLOB E_APO_CONFIG_FILES "config-files/*.txt")
file(COPY ${E_APO_CONFIG_FILES} DESTINATION ${CMAKE_CURRENT_BINARY_DIR}/config-files)
