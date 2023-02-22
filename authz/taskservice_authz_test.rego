package authz

test_post_allowed {
    allow with input as {
        "path": ["tasks"], 
        "method": "POST",
        "user": "miron",
        "owner": "miron"
    }
}

test_post_anonymous_denied {
    not allow with input as {
        "path": ["tasks"], 
        "method": "POST",
        "owner": "miron"
    }
}

test_post_another_user_denied {
    not allow with input as {
        "path": ["tasks"], 
        "method": "POST",
        "user": "johnsmith",
        "owner": "miron"
    }
}

test_post_admin_allowed {
    allow with input as {
        "path": ["tasks"], 
        "method": "POST",
        "user": "admin",
        "owner": "miron"
    }
}

test_get_with_id_allowed {
    allow with input as {
        "path": ["tasks", "42"],
        "taskid": "42",
        "method": "GET",
        "user": "miron",
        "owner": "miron"
    }
}

test_get_anonymous_denied {
    not allow with input as {
        "path": ["tasks", "42"],
        "taskid": "42",
        "method": "GET",
        "owner": "miron"
    }
}

test_get_another_user_denied {
    not allow with input as {
        "path": ["tasks", "42"],
        "taskid": "42", 
        "method": "GET",
        "user": "johnsmith",
        "owner": "miron"
    }
}

test_get_admin_allowed {
    allow with input as {
        "path": ["tasks", "42"],
        "taskid": "42", 
        "method": "GET",
        "user": "admin",
        "owner": "miron"
    }
}

test_get_list_with_id_allowed {
    allow with input as {
        "path": ["tasks"],
        "taskid": "42",
        "method": "GET",
        "user": "miron",
        "owner": "miron"
    }
}

test_get_list_anonymous_denied {
    not allow with input as {
        "path": ["tasks"],
        "taskid": "42",
        "method": "GET",
        "owner": "miron"
    }
}

test_get_list_another_user_denied {
    not allow with input as {
        "path": ["tasks"],
        "taskid": "42", 
        "method": "GET",
        "user": "johnsmith",
        "owner": "miron"
    }
}

test_get_list_admin_allowed {
    allow with input as {
        "path": ["tasks"],
        "taskid": "42", 
        "method": "GET",
        "user": "admin",
        "owner": "miron"
    }
}

test_delete_with_id_allowed {
    allow with input as {
        "path": ["tasks", "42"],
        "taskid": "42",
        "method": "DELETE",
        "user": "miron",
        "owner": "miron"
    }
}

test_delete_anonymous_denied {
    not allow with input as {
        "path": ["tasks", "42"],
        "taskid": "42",
        "method": "DELETE",
        "owner": "miron"
    }
}

test_delete_another_user_denied {
    not allow with input as {
        "path": ["tasks", "42"],
        "taskid": "42", 
        "method": "DELETE",
        "user": "johnsmith",
        "owner": "miron"
    }
}

test_delete_admin_allow {
    allow with input as {
        "path": ["tasks", "42"],
        "taskid": "42", 
        "method": "DELETE",
        "user": "admin",
        "owner": "miron"
    }
}

test_delete_all_denied {
    not allow with input as {
        "path": ["tasks"],
        "method": "DELETE",
        "user": "miron",
    }
}

test_delete_all_anonymous_denied {
    not allow with input as {
        "path": ["tasks"],
        "method": "DELETE",
    }
}

test_delete_all_admin_allow {
    allow with input as {
        "path": ["tasks"],
        "method": "DELETE",
        "user": "admin",
    }
}