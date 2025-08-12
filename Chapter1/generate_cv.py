from fpdf import FPDF

pdf = FPDF()
pdf.add_page()
pdf.set_font("Arial", size=12)

content = """Komartisin Lechisa Yabasa
Mechanical Engineer

Email: komartisinlechisa4@gmail.com | Phone: +251 973 369 489
Address: Addis Ababa, Ethiopia

Professional Summary
Mechanical Engineer with a BSc in Mechanical Engineering from Arbaminch University (2018) and over five years of experience in industrial machinery maintenance. Proven track record in diagnosing and repairing various mechanical systems and production machines. Worked with PVH Arvind Manufacturing PLC (May 2019 – Nov 2021) and currently employed at DLM PVT. LTD. CO. (Nov 2022 – Present). Completed a four-month internship at Ambo Mineral Water Share Company. Committed to delivering high-quality technical solutions in fast-paced industrial environments.

Work Experience
Mechanic
DLM PVT. LTD. CO. — Nov 29, 2022 - Present
- Maintaining and repairing both locally manufactured and imported machinery
- Diagnosing and resolving machine malfunctions efficiently
- Ensuring minimal downtime in production processes

Mechanic
PVH Arvind Manufacturing PLC — May 16, 2019 - Nov 25, 2021
- Worked on various industrial machines used in garment manufacturing
- Hands-on experience with lock stitch, chain stitch, overlock, buttonhole, embroidery, and pressing machines
- Became an expert in machine setup, operation, and maintenance

Intern – Mechanical Maintenance
Ambo Mineral Water Share Company — 4 Months
- Worked on washer, filler, labeler, boiler, compressor, and shrinkage machines
- Hands-on practice with maintenance, repair, and troubleshooting
- Participated in theoretical and practical training sessions

Education
BSc in Mechanical Engineering
Arba Minch Institute of Technology, Arba Minch University
2006 – 2010 (E.C.)

Preparatory School
Ambo Senior Secondary and Preparatory School
2004 – 2005 (E.C.)

Secondary School
Ambo Secondary School
2001 – 2003 (E.C.)

Primary School
Adis Ketema Primary School
1994 – 2000 (E.C.)

Technical Skills
- Mechanical Design & Analysis: SolidWorks, CATIA, AutoCAD
- Simulation & Analysis: MATLAB, ANSYS
- Documentation & Reporting: Microsoft Word, Excel
- Workshop Skills: Machine shop, welding, engine maintenance, tire and body shop
- Machinery Expertise: Embroidery machines, sleeve placket pressing machines, pocket attaching machines, conveyor systems

Projects & Practical Work
- BSc Thesis: Design and analysis of multi-directional screw jack
- Internship Project: Power screw conveyor design
- Practical training in mechanical labs: Engine lab, thermo-fluid lab, strength lab, welding shop, motor vehicle lab

Languages
- English: Reading (Excellent), Writing (Excellent), Listening (Excellent), Speaking (Very Good)
- Afaan Oromo: Excellent in all aspects
- Amharic: Excellent in all aspects

Hobbies & Interests
- Reading scientific and spiritual books
- Watching movies and listening to music
"""

pdf.multi_cell(0, 6, content)
pdf.output("Komartisin_Lechisa_Yabasa_CV_plain.pdf")
