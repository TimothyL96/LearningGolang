# **QuiLite**

---

### **Categorizing File**

- Per dataset is categorized to its own package
- Global libraries are at their own library, such as Key package in the key folder

---

### **Get Method**

- For all **get methods**, they do not have `Get` keyword in front.
- So a normal `GetTask()` will be named `Task()` instead


- For **get methods** that returns a relation, it will always return a pointer
- For **get methods** that return a normal field such as an int or string, it will return a value instead of pointer
---

### **Public vs Private members**

- In Golang, any fields, functions, struct or any other type: If the name starts with capital letter,  it will be accessible by other package.    
- Also known as **_exported type_**.


- Vice versa, if the name starts with small capital letter, it will not be accessible in other package.
- Also known as **_unexported type_**

---

### **Checking method receiver for nil**

- All **_get methods_** implementation have to check if the receiver is null.
    - For some of the methods, if the receiver is nil, always `return nil`
    - There's exception for get methods, which are get methods that return normal fields like int or string,
    see below _(Panic)_.
    
    
- For most of the implementation for **_unexported non-get methods_**, such as functions used for declarative propagation, there is a need to check
if the receiver is null.
    - If yes, then just simply `return` empty
    - Such scenarios usually happens when the propagator calls the method through chaining `A.B().C().D()`


- When using **exported non-get methods**, always check for nil **manually** as a caller.

    - This is because a valid method receiver is always required when using this kind of methods

        **_Example:_ To use A.B().ExportedNonGetMethod()**

            if A.B() != nil {
                A.B().ExportedNonGetMethod()       // Do something with it
            }
            
            
- **_Panic_** must always be used when something is not expected to work that way.
    - For example, when implementing **exported non-get methods**, call _panic_ if the method receiver is nil.
        - This is because the caller of such methods should be checking it themself.
       
    - Also, for **get methods** that returns a normal field such as an int or string, after checking
    that the method receiver is nil, call _panic_ as this is not expected behavior.
        - Similarly like above, the user has to check if the instance is nil or not before using its normal fields
  
        
- Any method with receiver but not used in the body can skip the checking for valid receiver
<br>

##### **Making sure all get methods follow the rules above will enable the chaining of multiple _pointer_ fields such as `A.B().C().D().E()` without having to check for nil**

---

### **Accessing fields**

- In get methods, directly access the fields to return the value

- Similarly in set methods, use direct access to set the fields
    - Exception is if the field has to be 'Get' before setting its value, then use the Get method implemented above
    
- All other cases, only use methods to access struct fields

---