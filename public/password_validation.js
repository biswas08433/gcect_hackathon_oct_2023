let password = $("#password");
let retype_password = $("#retype-password")
let lowercase_message = $("#letter");
let capital_message = $("#capital");
let number_message = $("#number");
let length_message = $("#length");

let handler = function () {
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

}

password.keyup(handler)
retype_password.keyup(handler)
