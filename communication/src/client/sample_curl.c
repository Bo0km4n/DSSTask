#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "curl/curl.h"

typedef struct {
    char *m;
    size_t size;
} Memory;

size_t write_callback(char *ptr, size_t size, size_t nmemb, void *userdata);

int main(int argc, char **argv)
{
    CURL *curl;
    CURLcode ret;
    Memory body;

    curl_global_init(CURL_GLOBAL_ALL);

    curl = curl_easy_init();
    if (curl == NULL) {
        fprintf(stderr, "Error: curl_easy_init() failed\n");
        return 1;
    }

    body.m = (char *)malloc(1);
    body.m[0] = '\0';
    body.size = 0;

    curl_easy_setopt(curl, CURLOPT_URL, "http://localhost:18888/api/v1/struct");
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &body);
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);

    ret = curl_easy_perform(curl);
    if (ret != CURLE_OK) {
        fprintf(stderr, "Error: curl_easy_perform() failed: %s\n", curl_easy_strerror(ret));
        return 1;
    }

    printf("reponse length: %zu\n", body.size);
    for (int i=0; i< body.size;i++) {
      printf("%c", body.m[i]);
    }

    curl_easy_cleanup(curl);

    curl_global_cleanup();

    return 0;
}

size_t write_callback(char *ptr, size_t size, size_t nmemb, void *userdata)
{
    char *p;
    Memory *mem = (Memory *)userdata;
    size_t append_size = size * nmemb;

    p = (char *)realloc(mem->m, mem->size + append_size + 1);
    if (p == NULL) {
        return 0;
    }
    memcpy(p + mem->size, ptr, append_size);

    mem->m = p;
    mem->size += append_size;
    mem->m[mem->size] = '\0';

    return size * nmemb;
}