import { Header } from "@/components/Header";
import { DepositsListCard } from "@/components/DepositsListCard";
import { PermitsListCard } from "@/components/PermitsListCard";

export default function DashboardPage() {
    return (
        <>
            <Header />
            <main className="flex flex-col gap-8 w-full max-w-[1400px] mx-auto px-[60px] pt-[120px] pb-[60px]">
                <div className="grid grid-cols-2 gap-8 items-start">
                    <DepositsListCard />
                    <PermitsListCard />
                </div>
            </main>
        </>
    );
}
