import {useContext, useEffect, useState} from "react"
import {EnvironmentContext} from "@kanban/environment"
import styles from "./styles.module.scss"

type Props = {
    onChange: (tags: string[]) => void
    selected?: string[]
}
export default function TagSelector({onChange, selected}: Props) {

    const {environment} = useContext(EnvironmentContext)

    const tags = environment?.tags

    const [newSelected, setNewSelected] = useState<string[]>(selected ?? [])
    const [firstRender, setFirstRender] = useState<boolean>(true)

    useEffect(() => {
        setNewSelected(selected ?? [])
    }, [selected])

    useEffect(() => {
        if (firstRender) {
            setFirstRender(false)
            return
        }
        onChange(newSelected)
    }, [newSelected])

    function checked(uuid: string) {
        return newSelected.includes(uuid)
    }

    function handleChange(uuid: string) {
        if (newSelected.includes(uuid)) {
            setNewSelected(newSelected.filter((t: string) => t !== uuid))
        } else {
            setNewSelected([...newSelected, uuid])
        }
    }

    return (
        <>
            <div className={styles.selected}>
                {newSelected?.map((uuid) => {
                    let tag = tags?.find((tag) => tag.uuid === uuid)
                    if (!tags) {
                        return <></>
                    }
                    return (
                        <div key={uuid} style={{
                            backgroundColor: `${tag?.color}90`,
                        }}>
                            {tag?.name}
                        </div>
                    )
                })}
            </div>
            {open ? <div className={styles.availableTags}>
                {tags?.map((tag) => (
                    <>
                        <label key={tag.uuid} className={styles.tag}
                               htmlFor={`tag-select-${tag.uuid}`}
                        >
                            <input id={`tag-select-${tag.uuid}`} type="checkbox" className={styles.checkbox}
                                   checked={checked(tag.uuid)}
                                   onChange={() => handleChange(tag.uuid)}/>

                            <span className={styles.name} style={{
                                backgroundColor: `${tag.color}90`,
                            }}>
                                {tag.name}
                            </span>
                        </label>

                    </>
                ))}

            </div>
        </>
    )
}