export default {
    Menus: {
        Dashboard: "Dashboard",
        System: "System",
        Roles: "Roles",
        Features: "Features",
        Menus: "Menus",
        Users: "Users",
        Listing: "Listing",
        NewUser: "New User"
    },
    Home: "Home",
    Search: "Search",
    AdvanceFilter: "Advance Filter",
    ActivateBtn: "Activate",
    DeactivateBtn: "Deactivate",
    ExportBtn: "Export",
    ImportBtn: "Import",
    MoreActions: "More Actions",
    BackBtn: "Back",
    EditBtn: "Edit",
    DeleteBtn: "Delete",
    Filter: "Filter",
    RememberFilters: "Remember Filters",
    SearchBtn: "Search",
    ResetBtn: "Reset",
    Yes: "Yes",
    No: "No",
    RowsPerPage: "Per Page",
    SaveBtn: "Save",
    CancelBtn: "Cancel",
    ValidationErrorsText: "There are some validation errors, please fix them first.",
    ViewBtn: "View",
    CreatedAt: "Created At",
    Active: "Active",
    Actions: "Actions",
    Delimeter: "Delimeter",
    TextEncloser: "Text Encloser",
    Caption: "Caption",
    Column: "Column",
    SelectColumns: "Select Columns",
    SaveExpBtn: "Save & Export",
    DownloadSampleBtn: "Download Sample",
    UploadFile: "Upload File",
    Format: "Format",
    PreviewData: "Preview Data",
    SelectFileMsg: "Please select a file",
    Roles: {
        Listing: {
            PageTitle: "Roles Management",
            NewBtn: "New Role",
            PermissionsBtn: "Permissions",
        },
        Breadcrumbs: {
            Roles: "Roles"
        },
        Fields: {
            Name: "Name"
        },
        Validations: {
            RoleRequired: "Please enter the role name.",
            RoleMaxLength: "Role name must not be longer than 20 characters.",
        }
    },
    Users: {
        MyProfile: {
            PageTitle: "My Profile",
            Tabs: {
                BasicInfo: "Basic Info",
                ChangePassword: "Change Password",
            }
        },
        Listing: {
            PageTitle: "Users Management",
            EditUser: "Edit User",
            NewBtn: "New User",
        },
        Breadcrumbs: {
            Users: "Users",
            Create: "Create",
            MyProfile: "My Profile",
        },
        Fields: {
            Name: "Name",
            Email: "Email",
            Roles: "Roles",
            Active: "Active",
            Password: "Password",
            ConfirmPassword: "Confirm Password",
            SelectProfilePicture: "Select Profile Picture",
            OldPassword: "Old Password",
        },
        Validations: {
            AllFieldsRequired: "All marked(*) fields are required.",
            PasswordValidationRules: "Password must have at least 1 digit,1 capital letter and 1 special characters(!,@,#,$,%,^,&,*). It must not start with special character.",
            ProfilePictureValidationRules: "Only png, jpg, gif format allowed for Profile Picture. Maximum 2mb file size is allowed.",
            NameRequired: "Please enter your name.",
            NameMaxLength: "Name must not be longer than 20 characters.",
            EmailRequired: "Please enter your email address.",
            EmailInvalid: "Please enter valid email address.",
            PasswordRequired: "Please enter your password.",
            PasswordText1: "Password must be atleast",
            PasswordText2: "characters long.",
            PasswordText3: "Password must not be greater than",
            PasswordText4: "characters.",
            PasswordNotMatch: "Password could not be matched.",
            RoleRequired: "Please select at least one role.",
        }
    },
    Posts: {
        Listing: {
            PageTitle: "Posts Management",
            EditPost: "Edit Post",
            NewBtn: "New Post",
        },
        Breadcrumbs: {
            Posts: "Posts",
            Create: "Create",
        },
        Fields: {
            Title: "Title",
            Description: "Description",
        },
        Validations: {

        }
    },
    Products: { 
        Listing: { 
            PageTitle: "Products Management", 
            EditProduct: "Edit Product", 
            NewBtn: "New Product", 
        }, 
        Breadcrumbs: { 
            Products: "Products", 
            Create: "Create", 
        }, 
        Fields: { 
            Name1: "Name1", 
            Name2: "Name2", 
            Price: "Price", 
            CategoryId: "Category id", 
        }, 
        Validations: {} 
    }
}