#include <stdio.h>
#include <stdlib.h>

#include "net.h"
#include "parser.h"
#include "descriptor.h"
#include "task.h"
#include "person.h"
#include "serializer.h"


struct person_t build_person_from_task(const struct task_t task) {
    struct person_t person = new_person();
    size_t person_name_len = strlen(task.str1) + strlen(task.str2) + 1;
    free(person.name);
    
    person.name = malloc(sizeof(char) * person_name_len);
    memset(person.name, '\0', person_name_len);
    strcat(person.name, task.str1);
    strcat(person.name, task.str2);
    return person;
}

int main(void) {
    Memory *buf = (Memory *)malloc(sizeof(Memory));
    buf->memory = NULL;
    buf->size = 0;
    
    fetch_struct(buf);

    printf("fetched data: \n");
    for (int i = 0; i < buf->size; i++) {
        printf("0x%02x, ", buf->memory[i] & 0x000000FF);
    }
    printf("\n");

    const unsigned char *bytes = (const unsigned char *)buf->memory;
    
    const size_t len = buf->size;
    instance_t inst = parse(bytes, len);
    
    struct task_t t = cast_task(inst);
    preview_task(t);
    
    struct person_t p = build_person_from_task(t);
    preview_person(p);

    instance_t p_inst = inter_serialize_person(p);
    struct bytes_t p_bytes = serialize(p_inst);

    post_struct(p_bytes);
}