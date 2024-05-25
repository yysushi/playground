# Play TextFSM

<https://github.com/google/textfsm/wiki/TextFSM>

- value definitions <- per line
- state definitions <- per block
  - state name
  - state rule definitions <- per line

- state rule definitions: `^regex [-> action]`
  - action: **A.B C** format
    - **A**: Line Actions, actions on the input line
      - **Next**: Finish with the input line, read in the next and start matching again from the start of the state.
      - **Continue**: Retain the current line and do not resume matching from the first rule of the state. Continue processing rules as if a match did not occur (value assignments still occur).
    - **B**: Record Actions, actions on the values collected so far
      - **NoRecord**: Do nothing.
      - **Record**: Record the values collected so far as a row in the return data. Non Filldown values are cleared. Note: No record will be output if there are any 'Required' values that are unassigned.
      - **Clear**: Clear non Filldown values.
      - **ClearAll**: Clear all values.
    - **C**: State transitions
    - the default implicit action is **Next.NoRecord**
- value definitions: *Value*` [option[,option...]] name regex`
  - option: Extra options regarding the value
    - **Filldown**: The previously matched value is retained for subsequent records (unless explicitly cleared or matched again). In other words, the most recently matched value is copied to newer rows unless matched again.
    - **Key**: Declares that the fields contents contribute to the unique identifier for a row.
    - **Required**: The record (row) is only saved into the table if this value is matched.
    - **List**: The value is a list, appended to on each match. Normally a match will overwrite any previous value in that row.
    - **Fillup**: Like Filldown, but populates upwards until it finds a non-empty entry. Not compatible with Required or List.
