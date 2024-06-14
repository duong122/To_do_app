import React from "react";
import "../App.css"
export const TodoForm = () => {
    return  (
        <form className="TodoForm">
            <input type="text" className='todo-input' placeholder="what is the task today"/>
            <button type="submit" className='todo-btn'> Add task </button>

        </form>
    );
}

