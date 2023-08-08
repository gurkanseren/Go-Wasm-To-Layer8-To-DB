const go = new Go();
WebAssembly.instantiateStreaming(fetch("module.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
  }
);

function displayImage(imageURL) {
  const imageElement = document.getElementById("imageElement");
  const scrollDownMsg = document.getElementById("scrollDownMsg");
  // Show the scroll down message
  scrollDownMsg.style.display = "block";
  imageElement.src = imageURL;
}

function showLoginForm() {
  const loginButton = document.getElementById("loginButton");
  const registerButton = document.getElementById("registerButton");
  const loginForm = document.getElementById("loginForm");
  const backButton = document.getElementById("backButton");
  const initialButtons = document.getElementById("initialButtons");

  if (loginButton && registerButton && loginForm && backButton) {
    loginButton.style.display = "none";
    registerButton.style.display = "none";
    loginForm.style.display = "block";
    backButton.style.display = "block";
    initialButtons.style.display = "none";
  }
}

function showRegisterForm() {
  const loginButton = document.getElementById("loginButton");
  const registerButton = document.getElementById("registerButton");
  const registerForm = document.getElementById("registerForm");
  const backButton = document.getElementById("backButton");
  const initialButtons = document.getElementById("initialButtons");

  if (loginButton && registerButton && registerForm && backButton) {
    loginButton.style.display = "none";
    registerButton.style.display = "none";
    registerForm.style.display = "block";
    backButton.style.display = "block";
    initialButtons.style.display = "none";
  }
}

function goBack() {
  const loginButton = document.getElementById("loginButton");
  const registerButton = document.getElementById("registerButton");
  const loginForm = document.getElementById("loginForm");
  const registerForm = document.getElementById("registerForm");
  const backButton = document.getElementById("backButton");
  const initialButtons = document.getElementById("initialButtons");
  const imageElement = document.getElementById("imageElement");
  const scrollDownMsg = document.getElementById("scrollDownMsg");
  const loginMessage = document.getElementById("loginMessage");
  const loginPassword = document.getElementById("loginPassword");
  const loginUsername = document.getElementById("loginUsername");
  const registerPassword = document.getElementById("registerPassword");
  const repeatPassword = document.getElementById("repeatPassword");
  const registerUsername = document.getElementById("registerUsername");

  if (
    loginButton &&
    registerButton &&
    loginForm &&
    registerForm &&
    backButton
  ) {
    loginButton.style.display = "block";
    registerButton.style.display = "block";
    loginForm.style.display = "none";
    registerForm.style.display = "none";
    backButton.style.display = "none";
    initialButtons.style.display = "block";
    imageElement.src = "";
    scrollDownMsg.style.display = "none";
    loginMessage.style.display = "none";
    loginPassword.value = "";
    loginUsername.value = "";
    registerPassword.value = "";
    registerUsername.value = "";
    repeatPassword.value = "";
  }
}

function regUser() {
  const username = document.getElementById("registerUsername").value;
  const password = document.getElementById("registerPassword").value;
  const repeatPassword = document.getElementById("repeatPassword").value;
  const repeatPasswordError = document.getElementById("repeatPasswordError");

  if (password !== repeatPassword) {
    repeatPasswordError.innerText = "Passwords do not match";
    repeatPasswordError.style.color = "red";
    return;
  }

  // Call your Go function to register the user
  registerUser(username, password);
}

function regUserError() {
  const repeatPasswordError = document.getElementById("repeatPasswordError");
  repeatPasswordError.innerText = "User already exists";
  repeatPasswordError.style.color = "red";
}

function regUserSuccess() {
  const repeatPasswordError = document.getElementById("repeatPasswordError");
  repeatPasswordError.innerText = "User registered";
  repeatPasswordError.style.color = "green";
}

function logUser() {
  const username = document.getElementById("loginUsername").value;
  const password = document.getElementById("loginPassword").value;
  const choice = document.getElementById("choice").value;

  // Call your Go function to login the user
  const loginResult = loginUser(username, password, choice);
}

function loginSuccess(imageURL) {
  const loginMessage = document.getElementById("loginMessage");
  loginMessage.innerText = "Login successful";
  loginMessage.style.color = "green";
  loginMessage.style.display = "block";
  loginMessage.style.textAlign = "center";
}

function loginError() {
  const loginMessage = document.getElementById("loginMessage");
  loginMessage.innerText = "Login failed";
  loginMessage.style.color = "red";
  loginMessage.style.display = "block";
  loginMessage.style.textAlign = "center";
}
