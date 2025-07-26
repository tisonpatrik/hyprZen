# HyprZen

## Installation Handling Strategy

### Human Interaction During Installation

When using `yay --noconfirm` for package installation, human interaction may still be required. Here's how we handle it:

#### Backend (Installation Service) Approach:

**1. Non-Blocking Installation**
- Run `yay --noconfirm` in a separate goroutine
- Use channels to communicate progress and status
- Don't block the UI while waiting for user input

**2. Input Detection & Pausing**
- Monitor the installation process output for prompts
- When we detect prompts like "Proceed with installation? [Y/n]" or "Enter password:", pause the installation
- Send a message to the frontend indicating user input is required

**3. State Management**
- Track installation state: `running`, `waiting_for_input`, `completed`, `failed`
- Store the current prompt/question that needs user response
- Maintain the installation process handle so we can resume it

**4. Graceful Handling**
- If user doesn't respond within a timeout, gracefully abort
- Provide clear error messages about what went wrong
- Allow user to retry or skip problematic packages

#### Frontend (UI) Approach:

**1. Dynamic UI States**
- **Normal Installation**: Show spinner + progress bar
- **Input Required**: Switch to input mode with the specific prompt
- **Error State**: Show error with retry/skip options
- **Paused State**: Show what's waiting and how to proceed

**2. User Input Handling**
- When backend signals input needed, show an input field
- Display the exact prompt from the installation process
- Send user response back to backend to resume installation
- Show a "waiting for input" indicator

**3. Error Recovery**
- If installation fails due to user input timeout, show options:
  - Retry the current package
  - Skip this package and continue
  - Abort entire installation
  - Show detailed error information

**4. Progress Persistence**
- Save installation progress so user can resume if interrupted
- Show which packages succeeded/failed/pending
- Allow selective retry of failed packages

#### Communication Flow:

```
Backend: Installation running → Detects prompt → Sends "input_required" message
Frontend: Receives message → Shows input UI → User responds → Sends response
Backend: Receives response → Resumes installation → Continues or fails
```

#### Safety Considerations:

1. **Timeout Protection**: Don't wait indefinitely for user input
2. **Process Management**: Properly handle subprocess lifecycle
3. **State Recovery**: Save progress so installation can be resumed
4. **Error Boundaries**: Don't let one failed package break everything
5. **User Control**: Always allow user to abort safely
