import {Button, Modal, Textarea, TextInput} from "@/components/ui"
import {FormEvent, useRef} from "react"
import styles from "./styles.module.scss"
import {toast} from "react-toastify"
import {IconX} from "@tabler/icons-react"
import {NameList} from "@kanban/types"

type props = {
    close: () => void,
    opened: boolean
    handleCreate: (workspace: string, name: string, description: string) => void
    workspace?: NameList
}


export default function CreateNewModal({opened, close, handleCreate, workspace}: props) {
    const nameRef = useRef<HTMLInputElement>(null)
    const descriptionRef = useRef<HTMLTextAreaElement>(null)

    if (!workspace) return null

    function handleSubmit(e: FormEvent<HTMLFormElement>) {
        if (!workspace) return

        e.preventDefault()
        const name = nameRef.current?.value
        const description = descriptionRef.current?.value
        if (name) {
            handleCreate(workspace.uuid, name, description ?? "")
        } else {
            toast("Name is required", {icon: <IconX/>, type: "error"})
        }

        close()
    }


    return (
        // @ts-ignore
        <Modal opened={opened} onClose={close} title={
            <>
                Create new Kanban Board in <i> {workspace.name} </i>
            </>
        }>
            <form onSubmit={handleSubmit} className={styles.createForm}>
                <TextInput ref={nameRef} label="Name" placeholder="Kanban Board Name" required/>
                <Textarea label="Description" autosize placeholder="Kanban Board Description"/>
                <Button variant="success" gradient type="submit">Create</Button>
            </form>
        </Modal>
    )
}