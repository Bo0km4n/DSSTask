#include <string.h>
#include "task.h"
#include "net.h"

struct task_t cast_task(struct inst instance) {
    struct task_t task = new_task();
    struct field_t *field = instance.u.object.clazz.field;
    struct classdata_t *classdata = instance.u.object.classdata;
    while (field != NULL && classdata != NULL) {
        if (strcmp("v", field->name) == 0) {
        task.v = classdata->i;
        } else if (strcmp("x", field->name) == 0) {
        task.x = classdata->b;
        } else if (strcmp("str1", field->name) == 0) {
        task.str1 = classdata->obj->u.str;
        } else if (strcmp("str2", field->name) == 0) {
        task.str2 = classdata->obj->u.str;
        }
        field = field->next;
        classdata = classdata->next;
    }
    return task;
};

struct task_t new_task() {
    struct task_t task;
    return task;
}

void preview_task(const struct task_t task) {
    printf("task.v: 0x%x\n", task.v);
    printf("task.x: %d\n", task.x);
    printf("task.str1: %s\n", task.str1);
    printf("task.str2: %s\n", task.str2);
}
