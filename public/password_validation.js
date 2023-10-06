let signup_form = $("#signup-form");
let password = $("#password");
let retype_password = $("#retype-password")
let lowercase_message = $("#letter");
let capital_message = $("#capital");
let number_message = $("#number");
let length_message = $("#length");

let password_validation_message = $("#password-validation-message");
let retype_password_message = $("#retype-password-message")

let submit_button = $("#submit-button")

let notification = $("#notification-1")

password_validation_message.hide()
retype_password_message.hide()
notification.hide()
submit_button.prop("disabled", true)


password.focus(function () {
    password_validation_message.show();
});
password.blur(function () {
    password_validation_message.hide();
});

retype_password.focus(function () {
    retype_password_message.show()
});

let password_validation_handler = function () {
    let success_class = "has-text-success";

    let lower_case = /[a-z]/g;
    let upper_case = /[A-Z]/g;
    let numbers = /[0-9]/g;
    let required_length = 8;

    // lowerkey validation
    if (password.val().match(lower_case)) {
        lowercase_message.addClass(success_class);
    } else {
        lowercase_message.removeClass(success_class);
    }

    if (password.val().match(upper_case)) {
        capital_message.addClass(success_class);
    } else {
        capital_message.removeClass(success_class);
    }

    if (password.val().match(numbers)) {
        number_message.addClass(success_class);
    } else {
        number_message.removeClass(success_class);
    }

    if (password.val().length >= required_length) {
        length_message.addClass(success_class);
    } else {
        length_message.removeClass(success_class);
    }
    if ($("#password").val() !== $("#retype-password").val()) {
        retype_password_message.removeClass("has-text-success");
        retype_password_message.addClass("has-text-danger");
        $("#retype-password-message-text").text("Password do not match!");
        submit_button.prop("disabled", true);
    } else {
        retype_password_message.removeClass("has-text-danger");
        retype_password_message.addClass("has-text-success");
        $("#retype-password-message-text").text("Password matches.");
        submit_button.prop("disabled", false);
    }

}

password.keyup(password_validation_handler);
retype_password.keyup(password_validation_handler);
