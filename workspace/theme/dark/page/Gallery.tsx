import Container from "../component/Container";

interface FriendLink {
    url: string
    name: string
    info?: string
}

interface Props {
    links: FriendLink[]
}

// @ts-ignore
import bblog from "bblog"
let params = bblog.getConfig();
let gallery = params.gallery

export default function Gallery() {
    return <Container>
        <h3>我的拍照、摄影，不过还在搭建中</h3>
        <div className="prose dark:prose-invert">

        </div>
    </Container>

}