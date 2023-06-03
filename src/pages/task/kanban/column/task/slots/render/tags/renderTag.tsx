import {TagsSlot} from "@kanban/column/task/slots/slotTypes"
import {kanbanTagType} from "@kanban/types"
import styles from "@kanban/column/task/slots/render/tags/styles.module.scss"


type Props = {
    tagSlot: TagsSlot
}

export default function RenderTags({tagSlot}: Props) {
    let tags = tagSlot.tag
    return (
        <div>
            {tags?.map((tag: kanbanTagType) => (
                <RenderTag key={tag.uuid} tag={tag}/>
            ))}
        </div>
    )
}

type RenderTagProps = {
    tag: kanbanTagType
}

export function RenderTag({tag}: RenderTagProps) {

    return (
        <div className={styles.tag} style={{backgroundColor: `${tag.color}80`, borderColor: tag.color}}>
            {tag.name}
        </div>
    )
}