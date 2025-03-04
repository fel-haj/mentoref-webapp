# MentoRef Web App

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![HTMX](https://img.shields.io/badge/HTMX-00ADD8?style=for-the-badge&logo=htmx&logoColor=white)](https://htmx.org/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white)](https://tailwindcss.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Hyperscript](https://img.shields.io/badge/Hyperscript-orange?style=for-the-badge&logo=javascript&logoColor=black)](https://hyperscript.org/)

MentoRef is a modern web application designed to connect job seekers and companies, offering features like "Blank Shot" applications, referrals, and mentorship opportunities.  It's built with a focus on simplicity and efficiency, leveraging Go for the backend, PostgreSQL for the database, and HTMX, _hyperscript and Tailwind CSS for a dynamic, responsive frontend.

## Features

*   **User Authentication:** Secure sign-up, sign-in, and sign-out functionality using JWT-based authentication.
*   **Dual User Roles:** Supports both individual users (job seekers) and company accounts.
*   **Dynamic Dashboard:**  A user-specific dashboard providing access to relevant features and information. (Currently under heavy development)
*   **Blank Shot Applications:**  Allows users to send out applications to a broad range of companies based on selected criteria.
*   **Referral System (Placeholder):**  Planned feature to facilitate job referrals.
*   **Mentorship Program (Placeholder):** Planned feature to connect users with mentors.
*   **Modern UI:**  A clean, responsive user interface built with Tailwind CSS and enhanced with HTMX and _hyperscript for dynamic interactions without extensive JavaScript.
* **View Transitions API:** Implements contemporary page transitions.

## Technologies Used

*   **Backend:** Go (with standard library `net/http`)
*   **Database:** PostgreSQL
*   **Frontend:**
    *   HTML Templates
    *   HTMX for dynamic content updates
    *   _hyperscript for client-side interactivity
    *   Tailwind CSS for styling
*   **Authentication:** JWT (JSON Web Tokens)
*   **Dependency Management:** Go Modules, npm (for Tailwind CSS and HTMX)

## Setup and Installation

**Prerequisites:**

*   Go (version 1.23 or later)
*   Node.js and npm (for Tailwind CSS and HTMX/Hyperscript)
*   PostgreSQL database
* A text editor or IDE

**Steps:**

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd <repository_name>
    ```

2.  **Install Dependencies:**

    *   Go dependencies are managed using Go Modules.
    *   Install npm dependencies:

        ```bash
        npm install
        ```

3.  **Environment Variables:**

    Create a `.env` file in the root directory of the project and set the following environment variables.  These are *crucial* for the application to function.

    ```
    DB_URL="postgres://user:password@host:port/database?sslmode=disable"
    SECRET_KEY="your-secret-key"  # A strong, randomly generated secret key
    PRIVATE_KEY="path/to/your/private-key.key" # Path to your private key for TLS
    CERTIFICATE="path/to/your/cert.crt"      # Path to your certificate file for TLS
    ```
    Replace the placeholders with your actual database connection string, a secret key, and the paths to your private key, and certificate.

4.  **Build Tailwind CSS:**

    ```bash
    npm run build
    ```

5.  **Build and Run the Go Application:**

    ```bash
    go build -o bin/app cmd/main.go
    ./bin/app
    ```

6.  **Access the Application:**

    The application will be accessible at `https://localhost:443`.  Note the `https` - the application is configured to run with TLS.

## Contact

For inquiries, please contact info@mentoref.com (placeholder email).
