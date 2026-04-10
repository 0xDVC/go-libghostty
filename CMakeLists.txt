cmake_minimum_required(VERSION 3.19)
project(go-libghostty LANGUAGES C)

include(FetchContent)
FetchContent_Declare(ghostty
    GIT_REPOSITORY https://github.com/ghostty-org/ghostty.git
    GIT_TAG 48a01b8bd51b0cf4ba3ed281a6662ae131ee8239
)
FetchContent_MakeAvailable(ghostty)
