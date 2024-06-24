import React from "react"
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons'
import { faTrash } from '@fortawesome/free-solid-svg-icons'

export const Todo = () => {
    return (
        <div>
            <p>This is job need to done</p>
            <FontAwesomeIcon className="edit-icon" icon={faPenToSquare}></FontAwesomeIcon>
            <FontAwesomeIcon className="delete-icon" icon={faTrash}/>
        </div>
    )
}