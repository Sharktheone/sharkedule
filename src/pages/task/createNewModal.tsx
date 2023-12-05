import {Button, Modal, Textarea, TextInput} from "@/components/ui"
import {FormEvent, useRef} from "react"
import {notifications} from "@mantine/notifications"

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
            notifications.show({title: "Error", message: "Name is required", color: "red"})
        }

        close()
    }

    return (
        <Modal opened={opened} onClose={close} title="Create new Kanban Board">
            <form onSubmit={handleSubmit}>
                <TextInput mb="lg" ref={nameRef} withAsterisk label="Name" placeholder="Kanban Board Name" required/>
                <Textarea mb="lg" label="Description" placeholder="Kanban Board Description"/>
                <Button variant="gradient" gradient={{from: "green", to: "lime"}} type="submit">Create</Button>
            </form>
        </Modal>
    )
}