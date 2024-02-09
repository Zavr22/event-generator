# Google Calendar Event Generator

This script is designed to automate the creation of events in Google Calendar. It efficiently generates and adds events to your Google Calendar based on user-defined parameters, such as the date range and the number of events per day. This tool is ideal for users looking to bulk populate their calendars with placeholder events, for testing purposes, or for setting up a series of reminders or activities over a specific period.

## Features

- **Custom Date Range**: Define a specific start and end date for event generation.
- **Events Per Day**: Specify the number of events to generate for each day within the date range.
- **Randomized Event Times**: Events are scheduled at random times within predefined time ranges to simulate a realistic calendar.
- **Simple Configuration**: Easy to configure with a JSON file or direct input.
- **Google Calendar API Integration**: Seamlessly integrates with Google Calendar API for adding events directly to your calendar.

## Prerequisites

Before you start using this script, make sure you have the following:

- **Google Cloud Platform Account**: A project set up on Google Cloud Platform with the Calendar API enabled.
- **API Credentials**: A `credentials.json` file obtained from your Google Cloud project that allows access to the Google Calendar API.
- **Google Calendar API Library**: Ensure the Google Calendar API Go client library is installed.

## Setup

1. **Enable the Google Calendar API**: Visit the [Google Cloud Console](https://console.cloud.google.com/), create a new project or select an existing one, search for the Google Calendar API, and enable it.

2. **Create Credentials**:
    - In the Google Cloud Console, navigate to the "Credentials" page.
    - Click "Create Credentials" and select "OAuth client ID".
    - If prompted, configure the OAuth consent screen.
    - For application type, select "Desktop app" and give it a name.
    - Download the generated credentials and save them as `credentials.json` in the project directory.

3. **Install Dependencies**: Run the following command to install the Google Calendar API client library for Go: 
`go get google.golang.org/api/calendar/v3`

## Configuration

Create a `input.json` file in the root directory of the project with the following structure:

```json
{
"startDate": "YYYY-MM-DD",
"endDate": "YYYY-MM-DD",
"eventsPerDay": 2
}
```

## Usage
1. **Run the Script**: Execute the script by running the following command in your terminal:
`go run main.go`

2. **Authorize Access**: On the first run, you'll be prompted to authorize access to your Google Calendar. Follow the instructions in the terminal to complete the authentication process.

3. **Check Your Calendar**: Once the script completes execution, check your Google Calendar to see the generated events.

## Customization

To customize the event generation further, such as adjusting the time ranges for event scheduling or modifying the event summary, edit the `generateEvents` function within the script.

## Troubleshooting

If you encounter issues with authentication or API limits, verify your Google Cloud project settings, check the API quotas, and ensure your OAuth consent screen is correctly configured.