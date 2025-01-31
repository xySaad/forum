
export const NewReference = (initialValue) => {
    let reference = initialValue;

    return (newValue) => {        
      if (newValue == undefined) {
        return reference
      }
  
      if (typeof newValue == "function") {
        reference = newValue(reference)
      } else {
        reference = newValue
      }
    }
  }
