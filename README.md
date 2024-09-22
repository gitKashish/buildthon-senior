### Project Explanation

This project involves creating a poll system using Svelte as the frontend and integrating it with an API from Kalp Studio’s gateway. Users can create polls, vote on them, and view results, with additional features like geolocation-based voting and poll expiration. The system will work in both online (API-enabled) and offline modes (local data fallback), and dummy records will simulate both expired and active polls. Here's a detailed breakdown of the project and its functionality:

---

### Key Features and Components:

1. **Home Page (Poll Fetching by ID)**:
   - **User Input for Poll ID**: The home page will allow users to enter a poll ID to fetch the poll details.
   - **API/Local Fallback**: If the API is working, it will fetch the poll data from the Kalp Studio API. If not, it will use dummy poll data stored locally. 
   - **Check for Poll Expiration**: Once the poll is fetched, the system checks whether the poll has expired by querying the API or local data.

2. **Poll Creation Page**:
   - **Create Poll Button**: A button on the home page that takes the user to the poll creation page.
   - **Poll Creation Form**: 
     - **Question Input**: An input field where the user can enter the poll question.
     - **Options Input**: A dynamically generated list where users can add multiple options for the poll.
     - **Location Input**: A list where the user can specify multiple locations (geographical areas or cities) related to the poll.
     - **Duration Input**: A field for setting the poll duration (in hours) after which the poll will expire.
   - **Submit Poll**: Once the form is filled out, the poll data will be submitted to the API to create the poll on the blockchain. If the API isn’t available, the poll will be stored locally.

3. **Poll Display (Active Poll)**:
   - **Displaying the Poll**: If the poll is not expired, the poll question, options, and location list will be displayed.
   - **Voting**: 
     - The user can select one of the poll options.
     - The user must also select one of the available locations from the list.
     - After selecting both the option and location, the user can submit their vote. 
   - **Vote Submission**: The vote will be submitted to the API or stored locally if the API is not available.

4. **Poll Display (Expired Poll)**:
   - **Showing Results**: If the poll has expired, instead of showing the voting options, the results will be displayed.
   - **Vote Breakdown**: The results will show each poll option with the number of people who chose that option.
   - **Location Breakdown**: A modal will display detailed information about which locations the voters were from when they cast their votes.
   
5. **Dummy Data for Offline Mode**:
   - **Dummy Active Poll**: The local environment will include sample data for an active poll, which can be fetched and displayed if the API is not available.
   - **Dummy Expired Poll**: Sample data for an expired poll will also be available locally, including the poll question, options, total votes per option, and voter locations.

6. **API Integration**:
   - The project uses Kalp Studio’s blockchain API for all operations related to creating polls, submitting votes, and checking poll states. The API endpoints used include:
     - **Create Poll**: To create new polls.
     - **Submit Vote**: To submit votes for a poll.
     - **Poll State**: To query the current state of a poll.
     - **Check if Poll is Open**: To determine whether a poll is still active or has expired.

7. **UI/UX Design**:
   - The user interface will be clean and user-friendly, with a focus on ease of navigation between pages (poll creation, poll voting, and viewing results).
   - **Modals**: Used for displaying additional information such as detailed vote breakdown by location.
   - **Responsive Design**: The page layouts will be responsive and mobile-friendly, ensuring that users on all devices can interact with the system smoothly.

---

### Workflow:

1. **Creating a Poll**:
   - The user clicks the "Create Poll" button and is taken to the poll creation form.
   - After filling in the question, options, locations, and duration, the poll is created via the API, or saved locally if the API is unavailable.

2. **Fetching and Voting on a Poll**:
   - The user enters a poll ID on the home page to fetch the poll details.
   - If the poll is still active, they can select an option, choose a location, and submit their vote.
   - If the poll is expired, the results are displayed, showing the options and vote count, with a breakdown of votes by location shown in a modal.

3. **Handling Expired Polls**:
   - When the user fetches a poll and it is expired, instead of voting options, the result summary is displayed.
   - The user can click on each option to view a detailed breakdown of where the votes came from (geographical locations).

4. **Offline Mode**:
   - If the API is unavailable, dummy data for both active and expired polls will be used, ensuring the system continues to function even without network connectivity.

---

### Local Environment Setup:
- **Offline Data**: Predefined polls (both active and expired) will be set up in local storage or within the Svelte app, simulating API responses.
- **Local API Fallback**: The system will automatically switch to using local data if the API is unreachable, ensuring seamless functionality in local environments.
- **API Key and Smart Contract Address**: The provided API key and smart contract address will be used for making calls to the Kalp Studio gateway, and the app will dynamically switch between API and local data as needed.

---

This setup allows the project to operate robustly in both online and offline modes, ensuring users can interact with polls, vote, and view results even if the API is temporarily unavailable.
