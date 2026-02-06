import { Header } from "@/components/Header";

export default function AboutPage() {
    return (
        <>
            <Header />
            <main className="flex flex-col gap-8 w-full max-w-[1000px] mx-auto px-[60px] pt-[120px] pb-[60px]">
                <div className="list-card about-card">
                    <div className="list-header" style={{ fontSize: '24px', marginBottom: '24px' }}>What is Terrace Finance?</div>

                    <div className="flex flex-col gap-6 text-[20px] leading-[1.6] opacity-90 font-[family-name:var(--font-source-sans)] text-justify">
                        <p>
                            Inspired by Subak, the ancient Balinese art of perfect distribution, we bring harmony to DeFi. Our platform is a non-custodial investment layer built for the Arc ecosystem and powered by Circle technologies.
                        </p>
                        <p>
                            Like water flowing through rice terraces, our system automatically reallocates stablecoins across lending protocols to capture the highest returns. The result is a seamless, automated flow of capital designed for peak efficiency and superior yield.
                        </p>
                    </div>
                </div>
            </main>
        </>
    );
}
