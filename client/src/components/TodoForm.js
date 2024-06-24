import React, { Component, useState } from "react";
import "../App.css"

export default class TodoForm extends Component {

    fetchList = () => {
        fetch('http://localhost:8888/api/todo/')
            .then((response) => response.json())
            .then(json => {
                // let posts = document.getElementById("post-list");
                // json.forEach(function (obj) {
                //     let li = document.createElement("li")
                //     li.appendChild(document.createTextNode(obj.title))
                //     posts.appendChild(li)
                // })
                console.log(json)
            })
        // console.log("I was click !!!")
    }

    render() {
        return (
            <div className="TodoForm">
                <input type="text" className='todo-input'
                    placeholder="what is the task today"
                    onChange={(e) => console.log(e.target.value)}
                />
                <button type="submit" className='todo-btn'> Add task </button>
                <ul id="post-list" className="Todo"></ul>
                <button onClick={this.fetchList} className="btn btn-primary">Fetch data</button>
            </div>
        )
    }

}

