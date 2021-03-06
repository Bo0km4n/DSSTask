#include <string.h>
#include "serializer.h"
#include "util.h"


void serialize_utf(struct blist *list, const char *str) {
    blist_push(list, bytes_from_string(str));
}

void serialize_classdata(struct blist *list, const struct object_t *object) {
    struct field_t *f = object->clazz.field;
    struct classdata_t *cd = object->classdata;

    while (f != NULL && cd != NULL) {
        switch (f->type) {
        case 'L':
        serialize_object(list, cd->obj);
        break;
        default:
        printf("not implemented (classdata type %c)\n", f->type);
        }

        f = f->next;
        cd = cd->next;
    }
}

void serialize_className(struct blist *list, const char *name) {
    serialize_utf(list, name);
}

void serialize_serialVersionUID(struct blist *list, unsigned long uid) {
    blist_push(list, bytes_from_long(uid));
}

void serialize_classDescFlags(struct blist *list, unsigned char flag) {
    blist_push(list, bytes_from_char(flag));
}

void serialize_fieldName(struct blist *list, const char *name) {
    serialize_utf(list, name);
}

void serialize_className1(struct blist *list, const struct inst *class_name) {
    serialize_object(list, class_name);
}

void serialize_fieldDesc(struct blist *list, const struct field_t *field) {
    blist_push(list, bytes_from_char(field->type));
    switch (field->type) {
    case 'L':	// object
        serialize_fieldName(list, field->name);
        serialize_className1(list, field->class_name);
        break;
    }

}

void serialize_fields(struct blist *list, const struct class_t *clazz) {
    unsigned short fields_len = 0;
    struct field_t *f = clazz->field;
    while (f != NULL) {
        fields_len++;
        f = f->next;
    }
    blist_push(list, bytes_from_short(fields_len));

    f = clazz->field;
    for (size_t i = 0; i < fields_len; i++) {
        serialize_fieldDesc(list, f);
        f = f->next;
    }
}

void serialize_endBlockData(struct blist *list) {
    blist_push(list, bytes_from_char(TC_ENDBLOCKDATA));
}

void serialize_classAnnotation(struct blist *list) {
    serialize_endBlockData(list);
}

void serialize_superClassDesc(struct blist *list, const struct class_t *clazz) {
    serialize_classDesc(list, clazz);
}

void serialize_classDescInfo(struct blist *list, const struct class_t *clazz) {
    serialize_classDescFlags(list, clazz->flag);
    serialize_fields(list, clazz);
    serialize_classAnnotation(list);
    serialize_superClassDesc(list, NULL);
}

void serialize_newClassDesc(struct blist *list, const struct class_t *clazz) {
    blist_push(list, bytes_from_char(TC_CLASSDESC));
    serialize_className(list, clazz->name);
    serialize_serialVersionUID(list, clazz->uid);
    serialize_classDescInfo(list, clazz);
}

void serialize_nullReference(struct blist *list) {
    blist_push(list, bytes_from_char(TC_NULL));
}

void serialize_classDesc(struct blist *list, const struct class_t *clazz) {
    if (clazz == NULL) {
        serialize_nullReference(list);
    } else {
        serialize_newClassDesc(list, clazz);
    }
}

void serialize_newObject(struct blist *list, const struct inst *instance) {
    blist_push(list, bytes_from_char(TC_OBJECT));
    serialize_classDesc(list, &instance->u.object.clazz);
    serialize_classdata(list, &instance->u.object);
}

void serialize_newString(struct blist *list, const struct inst *instance) {
    blist_push(list, bytes_from_char(instance->type));

    switch (instance->type) {
    case TC_STRING:
        serialize_utf(list, instance->u.str);
        break;
    case TC_LONGSTRING:
        printf("not implemented (newString TC_LONGSTRING)\n");
        break;
    }
}

void serialize_object(struct blist *list, const struct inst *instance) {
    switch (instance->type) {
    case TC_OBJECT:
        serialize_newObject(list, instance);
        break;
    case TC_STRING:
        serialize_newString(list, instance);
        break;
    default:
        printf("not implemented (object: 0x%02x)\n", instance->type);
    }
}

struct bytes_t serialize(struct inst instance) {
    struct blist list = new_blist();

    blist_push(&list, bytes_from_short(STREAM_MAGIC));
    blist_push(&list, bytes_from_short(STREAM_VERSION));

    serialize_object(&list, &instance);

    struct bytes_t ret = blist_concat(list);
    blist_free(list);

    hexdump("serialize_result", ret.head, ret.len);
    return ret;
}
