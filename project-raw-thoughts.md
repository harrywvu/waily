
This is basically a glorified to-do app but with a twist. It's still a CRUD app but differently implemented.

V1. The Prototype w/ TUI

    ▗▖ ▗▖ ▗▄▖ ▗▄▄▄▖▗▖ ▗▖  ▗▖
    ▐▌ ▐▌▐▌ ▐▌  █  ▐▌  ▝▚▞▘ 
    ▐▌ ▐▌▐▛▀▜▌  █  ▐▌   ▐▌  
    ▐▙█▟▌▐▌ ▐▌▗▄█▄▖▐▙▄▄▖▐▌  
        - Date Today -
    x Wails added today! / Wail Stream is empty :(
        Wail Added! :D // when a new wail is added Wail Edited! // when an existing wail is edited
1. Add Daily Wail!
    Daily Wail: *Enter Win* // When enter is pressed, it gets added to the stream
    // back to main menu // 
2. View Wail Streams!
        Date                    
    Month Day, Year  < // Cursor for user to select // DEL button - Delete entire wail stream | INS button, EDIT mode wail stream
    Month Day, Year
    Month Day, Year

        Wail Stream
    - Month Day, Year -
Went to the gym and broke a PR
Stayed on calorie deficit
Read 18 pages of my novel   

        Wail Stream EDIT MODE
    - Month Day, Year -
Went to the gym and broke a PR  < // Cursor for user to select // DEL button - Delete wail from stream | INS button, Edit mode wail
Stayed on calorie deficit 
Read 18 pages of my novel   

    EDIT MODE WAIL
    Edit Wail: *Enter New Wail* // When enter is pressed, it updates the current wail
    // back to main menu // 


## **V1 – Daily Wails Prototype (CLI/TUI)**

### **Main Menu**

* [X] Show banner / title
* [X] Display today’s date
* [ ] Display Wails added today or message if empty
* [ ] Show success messages:

  * [ ] “Wail Added!” when a new wail is added
  * [ ] “Wail Edited!” when a wail is updated

---

### **1. Add Daily Wail**

* [ ] Prompt user: `Daily Wail: *Enter Win*`
* [ ] Add new Wail to today’s stream on Enter
* [ ] Return to main menu

---

### **2. View Wail Streams**

* [ ] List all available Wail Streams by date

  * [ ] Format: `Month Day, Year`
  * [ ] Cursor for user selection
* [ ] Key actions in stream list:

  * [ ] DEL → Delete entire Wail Stream
  * [ ] INS → Enter EDIT mode for selected Wail Stream

---

### **3. Wail Stream View**

* [ ] Display all Wails for selected stream
* [ ] Show date header: `- Month Day, Year -`
* [ ] List Wails in order added

---

### **4. Wail Stream EDIT MODE**

* [ ] Cursor to select individual Wail
* [ ] Key actions:

  * [ ] DEL → Delete selected Wail
  * [ ] INS → Edit selected Wail

---

### **5. EDIT MODE WAIL**

* [ ] Prompt user: `Edit Wail: *Enter New Wail*`
* [ ] Update selected Wail on Enter
* [ ] Return to main menu

---

### **General**

* [ ] All streams stored in JSON
* [ ] Basic CRUD implemented:

  * [ ] Create (Add Wail)
  * [ ] Read (View Stream)
  * [ ] Update (Edit Wail)
  * [ ] Delete (Wail / Stream)