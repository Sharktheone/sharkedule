import {Text} from "@mantine/core"
import {useStyles} from "./styles"
import {kanbanTaskType} from "../../types"
import {useState} from "react"
import styles from "./styles.module.scss"
import {IconCircleCheck, IconTrash, IconX} from "@tabler/icons-react"
import {api} from "@/api/api"
import {notifications} from "@mantine/notifications"
import {useNavigate} from "react-router-dom"

type TaskProps = {
    task: kanbanTaskType
    renameTask: (uuid: string, name: string) => void
    boardUUID: string
    columnUUID: string
}


export default function Task({task, renameTask, boardUUID, columnUUID}: TaskProps) {
    const {classes, cx} = useStyles()
    const [editable, setEditable] = useState(false)
    const navigate = useNavigate()

    function editText() {
        setEditable(true)
    }

    function handleBlur(e: any) {
        setEditable(false)
        renameTask(task.uuid, e.target.innerText)
    }

    function handleDelete() {
        api.delete(`/kanbanboard/${boardUUID}/column/${columnUUID}/task/${task.uuid}/delete`).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: "res.data", color: "red", icon: <IconX/>})
                } else {
                    notifications.show({title: "Success", message: "Deleted Task", color: "green"})
                    navigate("#")

                }

            }
        ).catch(
            (err) => {
                notifications.show({title: "Error", message: err.message, color: "red", icon: <IconX/>})
                console.log(err)
            }
        )
    }

    return (
        <div className={`${cx(classes.task)} ${styles.task}`}>
            <div className={styles.name}>
                <IconCircleCheck/>
                <Text align="start" onClick={editText} onBlur={handleBlur} contentEditable={editable}>
                    {task.name}
                </Text>
            </div>
            <div className={styles.hover}>
                <div>
                    <button onClick={handleDelete}>
                        <IconTrash/>
                    </button>
                </div>
            </div>
        </div>
    )
}