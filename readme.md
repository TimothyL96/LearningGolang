# **Qui Lite**

---

### **Categorizing File**

- Per dataset is categorized to its own package
- Global libraries are at their own library, such as Key package in the key folder

---

### **Get Method**

- For all Get methods, they do not have Get keyword in front.
- So a normal `GetTask()` will be named `Task()` instead

---

### **Public vs Private members**

- In Golang, any fields, functions, struct or any other type: If the name starts with capital letter,  it will be accessible by other package.    
- Also known as **_exported type_**.


- Vice versa, if the name starts with small capital letter, it will not be accessible in other package.
- Also known as **_unexported type_**

---

### **Checking method receiver for nil**

- All **_get methods_** implementation have to check if the receiver is null.
    - If the receiver is nil, always `return nil`
    
    
- For most of the implementation for **_unexported non-get methods_**, such as functions used for declarative propagation, there is a need to check
if the receiver is null.
    - If yes, then just simply `return` empty


- When using **exported non-get methods**, always check for nil as a caller.

    - This is because a valid method receiver is always required when using this kind of methods

        **_Example:_ To use A.B().C()**

            if A.B() != nil {
                A.B().ExportedNonGetMethod()       // Do something with it
            }
            
- **_Panic_** must always be used when something is not expected to work that way.
    - For example, when implementing **exported non-get methods**, panic if method receiver is nil.
        - This is because the caller of such methods should be checking it themself.
  
        
- Any method with receiver but not used in the body can skip the checking for valid receiver
<br>

##### **Making sure all get methods follow the rules above will enable the chaining of multiple pointers such as `A.B().C().D().E()` without having to check for nil**

---
