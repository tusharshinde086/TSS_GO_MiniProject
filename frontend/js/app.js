document.addEventListener("DOMContentLoaded", () => {
    const loginForm = document.getElementById("login-form");
    const registerForm = document.getElementById("register-form");
    const projectList = document.getElementById("project-list");
    const addProjectBtn = document.getElementById("add-project-btn");

    // Dummy data for projects (for demonstration purposes)
    const projects = [
        { id: 1, name: "Project Alpha" },
        { id: 2, name: "Project Beta" },
    ];

    // Function to display projects
    function displayProjects() {
        projectList.innerHTML = ""; // Clear existing projects
        projects.forEach(project => {
            const projectItem = document.createElement("div");
            projectItem.textContent = project.name;
            projectList.appendChild(projectItem);
        });
    }

    // Handle login form submission
    loginForm.addEventListener("submit", async (event) => {
        event.preventDefault(); // Prevent default form submission

        const email = document.getElementById("login-email").value;
        const password = document.getElementById("login-password").value;

        const response = await fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password }),
        });

        if (response.ok) {
            const user = await response.json();
            alert(`Welcome, ${user.name}!`);
            // Optionally, redirect to the projects page or update the UI
        } else {
            alert("Login failed. Please check your credentials.");
        }
    });

    // Handle registration form submission
    registerForm.addEventListener("submit", async (event) => {
        event.preventDefault(); // Prevent default form submission

        const name = document.getElementById("register-name").value;
        const email = document.getElementById("register-email").value;
        const password = document.getElementById("register-password").value;

        const response = await fetch("/api/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ name, email, password }),
        });

        if (response.ok) {
            const newUser  = await response.json();
            alert(`Registration successful! Welcome, ${newUser .name}!`);
            // Optionally, redirect to the login page or update the UI
        } else {
            alert("Registration failed. Please try again.");
        }
    });

    // Handle adding a new project (dummy implementation)
    addProjectBtn.addEventListener("click", () => {
        const projectName = prompt("Enter project name:");
        if (projectName) {
            const newProject = { id: projects.length + 1, name: projectName };
            projects.push(newProject);
            displayProjects();
        }
    });

    // Initial display of projects
    displayProjects();
});