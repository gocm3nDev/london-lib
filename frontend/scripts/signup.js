const username = document.getElementById('register-name');
const email = document.getElementById('register-email');
const password = document.getElementById('register-password');
const password_a = document.getElementById('register-password-a'); // password again
const register_button = document.getElementById('register-button');
const pass_strength = document.getElementById('pass-strength');

const inputs = [username, email, password, password_a];
const numbers = ['0','1','2','3','4','5','6','7','8','9'];
const specials = ['!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+'];

password.addEventListener("input", (e) => {
    const input = e.target.value;
    let numberCount = 0;
    let specialCount = 0;

    for (let char of input) {
        if (numbers.includes(char)) numberCount++;
        if (specials.includes(char)) specialCount++;
    }

    let message = "";

    if (input.length < 8) {
        message = "Minimum 8 characters required";
    } else if (numberCount < 2 || specialCount < 2) {
        message = "Password must contain at least 2 numbers and 2 special characters";
    } else if (input.length >= 8 && numberCount >= 2 && specialCount >= 2 && input.length < 10) {
        message = "<p>Password strength: <span class='minimum'>minimum requirements</span></p>";
    } else if (input.length >= 10 && numberCount >= 3 && specialCount >= 3) {
        message = "<p>Password strength: <span class='good'>good</span></p>";
    } else if (input.length >= 12 && numberCount >= 4 && specialCount >= 4) {
        message = "<p>Password strength: <span class='strong'>strong</span></p>";
    }

    pass_strength.innerHTML = message;
});



register_button.addEventListener("click", () => {
    if (inputs.some(el => !(el?.value || '').trim())) {
        alert("Please fill all fields");
        return;
    } else {
        if (password.value === password_a.value) {
            const now = new Date().toISOString();
            const user = {
                Name: username.value,     
                Email: email.value,
                PasswordHash: password.value,  
                IsActive: true,              
                CreatedAt: now,
            };

            fetch("http://localhost:8080/users", {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(user)
            })
                .then((response) => response.json())
                .then((data) => console.log(data))
                .catch((err) => console.log(err));
        } else {
            alert("Passwords are not match!");
            return;
        }
    }
});