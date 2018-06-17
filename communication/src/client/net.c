#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <curl/curl.h>
#include "net.h"

size_t write_memory_callback(void* ptr, size_t size, size_t nmemb, void* data) {
    if (size * nmemb == 0)
        return 0;

    size_t realsize = size * nmemb;
    Memory* mem = (Memory*)data;
    
    mem->memory = (char*)realloc(mem->memory, mem->size + realsize + 1);
    if (mem->memory) {
        memcpy(&(mem->memory[mem->size]), ptr, realsize);
        mem->size += realsize;
        mem->memory[mem->size] = 0;
    }
    return realsize;
}


void fetch_struct(Memory *buf) {
    CURL *curl = curl_easy_init();

    Memory* mem = malloc(sizeof(Memory*));

    mem->size = 0;
    mem->memory = NULL;

    char host[256];
    sprintf(host, "%s:%d/api/v1/struct", HOST, PORT);
    curl_easy_setopt(curl, CURLOPT_URL, host);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void *)mem);
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_memory_callback);
    curl_easy_perform(curl);
    curl_easy_cleanup(curl);
    buf->memory = realloc(buf->memory, mem->size + 1);
    buf->size = mem->size;
    for (int i = 0; i < mem->size; i++) {
        buf->memory[i] = mem->memory[i];
    }

    free(mem->memory);
    free(mem);

}
