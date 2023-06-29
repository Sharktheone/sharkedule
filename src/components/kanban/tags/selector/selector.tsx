import {useContext, useEffect, useState} from "react"
import {EnvironmentContext} from "@kanban/environment"
import styles from "./styles.module.scss"

type Props = {
    onChange: (tags: string[]) => void
}
export default function TagSelector({onChange}: Props) {

    const {environment} = useContext(EnvironmentContext)

    const tags = environment?.tags

    const [selected, setSelected] = useState<string[]>([])

    useEffect(() => {
        onChange(selected)
    }, [selected])

    function checked(uuid: string) {
        return selected.includes(uuid)
    }

    function handleChange(uuid: string) {
        if (selected.includes(uuid)) {
            setSelected(selected.filter((t: string) => t !== uuid))
        } else {
            setSelected([...selected, uuid])
        }
    }

    return (
        <>
            <div className={styles.avaiableTags}>
                {tags?.map((tag) => (
                    <>
                        <label key={tag.uuid} className={styles.tag}
                        htmlFor={`tag-select-${tag.uuid}`}
                        >
                            <input id={`tag-select-${tag.uuid}`} type="checkbox" className={styles.checkbox} checked={checked(tag.uuid)}
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