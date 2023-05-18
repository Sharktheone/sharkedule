import {Button, Modal, TextInput} from "@mantine/core";
import {useRef} from "react";


type props = {
    close: () => void
    opened: boolean
    addTask: (name: string) => void
}

export default function NewTaskModal({close, opened, addTask}: props) {

    const inputRef = useRef<HTMLInputElement>(null)

    function handleSubmit(e: any) {
        e.preventDefault()
        const name = inputRef.current?.value
        if (name) {
            addTask(name)
        }
        close()
    }

    return (
        <Modal opened={opened} onClose={close} title={"Add new task"}>
            <form onSubmit={handleSubmit}>
                <TextInput mb="lg" ref={inputRef} withAsterisk label="Name" placeholder="Task name" required/>
                <Button variant="gradient" gradient={{from: "green", to: "lime"}} type="submit">Create</Button>
            </form>

        </Modal>
    )
}