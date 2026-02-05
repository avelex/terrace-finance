import { Header } from "@/components/Header";
import { OperationsTable } from "@/components/OperationsTable";

export default function ExplorerPage() {
    return (
        <>
            <Header />
            <main className="main-container">
                <OperationsTable />
            </main>
        </>
    );
}
