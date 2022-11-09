export const $ApiBaseUrl = 'http://api.govueadmin.localhost';

export const URLs = {
    REG_ROLES: $ApiBaseUrl + '/api/auth/reg/roles',
    POST_LOGIN: $ApiBaseUrl + '/users/auth/login',
    FORGOT_PWD: $ApiBaseUrl + '/forgot-password',
    RESET_PWD: $ApiBaseUrl + '/reset-password',
    ADMIN_PANEL: $ApiBaseUrl + '/admin/dashboard',
    EXPORT_UNLINK: $ApiBaseUrl + '/api/export/unlink',
    USER_LIST: $ApiBaseUrl + '/api/users',
    USER_COOKIES: $ApiBaseUrl + '/api/users/cookies',
    USER_ACCESS: $ApiBaseUrl + '/api/users/checkAccess',
    USER_PICTURE_REMOVE: $ApiBaseUrl + '/api/users/removePic',
    USER_EXPORT_SETTINGS: $ApiBaseUrl + '/api/users/export/settings',
    USER_EXPORT: $ApiBaseUrl + '/api/users/export',
    USER_IMPORT_SETTINGS: $ApiBaseUrl + '/api/users/import/settings',
    USER_IMPORT: $ApiBaseUrl + '/api/users/import',
    ROLES_LIST: $ApiBaseUrl + '/api/roles',
    USER_REGISTER: $ApiBaseUrl + '/users/register',
    MY_USER: $ApiBaseUrl + '/users/me',
    ROLES_COOKIES: $ApiBaseUrl + '/api/roles/cookies',
    ROLES_PERMISSIONS: $ApiBaseUrl + '/api/roles/permissions',
    ROLE_ACCESS: $ApiBaseUrl + '/api/roles/checkAccess',
    FEATURE_LIST: $ApiBaseUrl + '/api/features',
    FEATURE_COOKIES: $ApiBaseUrl + '/api/features/cookies',
    FEATURE_ACCESS: $ApiBaseUrl + '/api/features/checkAccess',
    USER_ACT_DEACT: $ApiBaseUrl + '/api/users/actDeact',
    ROLE_ACT_DEACT: $ApiBaseUrl + '/api/roles/actDeact',
    USER_LOG_OUT: $ApiBaseUrl + '/users/auth/logout',
    MY_PROFILE: $ApiBaseUrl + '/api/users/profile',
    MENUS_LIST: $ApiBaseUrl + '/api/menus',
    MENUS_REORDER: $ApiBaseUrl + '/api/menus/reorder',
    MENUS_ACCESS: $ApiBaseUrl + '/api/menus/checkAccess',
    CRUD_MAKER: $ApiBaseUrl + '/api/crud/generate',
    LANG_SWITCH: $ApiBaseUrl + '/api/switchLang',
    SANCTUM_CSRF: $ApiBaseUrl + '/sanctum/csrf-cookie',
    PRODUCTS_LIST : $ApiBaseUrl + '/api/products',
    PRODUCTS_COOKIES : $ApiBaseUrl + '/api/products/cookies',
    PRODUCTS_ACCESS : $ApiBaseUrl + '/api/products/checkAccess',
    PRODUCTS_ACT_DEACT : $ApiBaseUrl + '/api/products/actDeact',
    PRODUCTS_EXPORT_SETTINGS : $ApiBaseUrl + '/api/products/export/settings',
    PRODUCTS_EXPORT : $ApiBaseUrl + '/api/products/export',
    PRODUCTS_IMPORT_SETTINGS : $ApiBaseUrl + '/api/products/import/settings',
    PRODUCTS_IMPORT : $ApiBaseUrl + '/api/products/import',
}