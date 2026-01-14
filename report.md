# Bug Report Log

## Entry 1

**Timestamp:** 0:01  
**Position:** Dashboard page  
**Problem:** Incorrect alignment of metric cards  
**Suggestion:** Arrange metric cards in a single grid occupying 2 rows, with feeds rendered after the metric cards.

## Entry 2

**Timestamp:** 0:14  
**Position:** Logical devices page  
**Problem:** When navigating to the page, the spinner is positioned at the top and is almost invisible  
**Suggestion:** Center the spinner vertically and horizontally in the middle of the screen so users can see that a request is being made to the server.

## Entry 3

**Timestamp:** 0:55  
**Position:** Edit project modal / dialog  
**Problem:** The date pickers were prefilled with the raw date string, which is not readable for an average user  
**Suggestion:** A properly formatted date should be displayed in the date picker components.

## Entry 4

**Timestamp:** 3:10  
**Position:** Organization table  
**Problem:** Some of the values displayed under the name column are not properly formatted  
**Suggestion:** For example, values that appear like “project_demo_ca” can be formatted properly. Packages like lodash's startCase method can help, or JavaScript's string replace method can be used.

## Entry 5

**Timestamp:** 3:35  
**Position:** Assign device to organization modal / dialog  
**Problem:** The popover component of the date picker doesn't close after a user selects a date; the user has to click outside to close it  
**Suggestion:** Use a state to track the opening and closing of the date picker's popover. Set the state to false after a user selects a date to close the popover. Refer to the ShadCN documentation for best practices.

## Entry 6

**Timestamp:** 5:12  
**Position:** Pedestrian count  
**Problem:** When a user selects a week from the select component, the text under the component reads “Week: 01 Jan 2025 to 04 Jan 2025”, for instance  
**Suggestion:** The selected option could be rendered in the text such that it appears like “{selected week}: 01 Jan 2025 to 04 Jan 2025”. A state can be used to dynamically render the selected week in the text.

## Entry 7

**Timestamp:** 6:59  
**Position:** Dashboard Layout  
**Problem:** Inconsistent layout  
**Suggestion:** A global layout component can be designed and used across all pages to maintain uniformity in margins and padding across the pages.

## Entry 8

**Timestamp:** 7:06  
**Position:** Logout button  
**Problem:** When a user clicks logout, they are redirected to the register page  
**Suggestion:** When a user clicks logout, they should be redirected to the login page instead of the register page.

## Entry 9

**Timestamp:** 7:30  
**Position:** Top Navbar  
**Problem:** Top navbar looks tiny; texts inside it are shrunk and almost invisible  
**Suggestion:** Perhaps a little extra padding can be added to make the texts, like user initials, more visible.

## Entry 10

**Timestamp:** 8:03  
**Position:** Top Navbar  
**Problem:** Inconsistent display of titles across pages. When navigating to a page like “Organisation Settings”, the title is boldly displayed on the navbar, whereas when navigating to the “Users” page, the title is not displayed  
**Suggestion:** Consistent display of page titles should be maintained across the app. The useLocation hook can help dynamically display titles when the user navigates in the app.

---
