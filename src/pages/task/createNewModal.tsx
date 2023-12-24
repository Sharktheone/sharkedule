import {Button, Modal, Textarea, TextInput} from "@/components/ui"
import {FormEvent, useRef} from "react"
import styles from "./styles.module.scss"
import {toast} from "react-toastify"
import {IconX} from "@tabler/icons-react"

type props = {
    close: () => void,
    opened: boolean
    handleCreate: (name: string, description: string) => void
}


export default function CreateNewModal({opened, close, handleCreate}: props) {
    const nameRef = useRef<HTMLInputElement>(null)
    const descriptionRef = useRef<HTMLTextAreaElement>(null)

    function handleSubmit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault()
        const name = nameRef.current?.value
        const description = descriptionRef.current?.value
        if (name) {
            handleCreate(name, description ?? "")
        } else {
            toast("Name is required", {icon: <IconX/>, type: "error"})
        }

        close()
    }

    return (
        <Modal opened={opened} onClose={close} title="Create new Kanban Board">
            <form onSubmit={handleSubmit} className={styles.createForm}>
                <TextInput ref={nameRef} label="Name" placeholder="Kanban Board Name" required/>
                <Textarea label="Description" placeholder="Kanban Board Description"/>
                <Button variant="success" gradient type="submit">Create</Button>
            </form>
        </Modal>
    )
}