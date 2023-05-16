import {Button, Modal, TextInput} from "@mantine/core"
import {FormEvent, useRef} from "react"
import {notifications} from "@mantine/notifications"

type props = {
    close: () => void,
    opened: boolean
    handleCreate: (name: string) => void
}


export default function CreateNewModal({opened, close, handleCreate}: props) {
    const inputRef = useRef<HTMLInputElement>(null)

    function handleSubmit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault()
        const name = inputRef.current?.value
        if (name) {
            handleCreate(name)
        } else {
            notifications.show({title: "Error", message: "Name is required", color: "red"})
        }

        close()
    }

    return (
        <Modal opened={opened} onClose={close} title="Create new Kanban Board">
            <form onSubmit={handleSubmit}>
                <TextInput mb="lg" ref={inputRef} withAsterisk label="Name" placeholder="Board Name" required/>
                <Button variant="gradient" gradient={{from: "green", to: "lime"}} type="submit">Create</Button>
            </form>

        </Modal>
    )
}