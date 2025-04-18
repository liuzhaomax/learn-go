-- Insert the primary category 'Physics'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Physics', 'Covers the broad field of Physics, including theoretical and experimental studies.', NULL);
-- Insert the primary category 'Mathematics'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Mathematics', 'Covers a wide range of mathematical fields including pure and applied mathematics.', NULL);
-- Insert the primary category 'Computer Science'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computer Science', 'Covers a wide range of computing topics including artificial intelligence, systems, and algorithms.', NULL);
-- Insert the primary category 'Quantitative Biology'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Quantitative Biology', 'Covers research in the quantitative analysis of biological systems, including mathematical modeling, computational biology, and systems biology.', NULL);
-- Insert the primary category 'Quantitative Finance'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Quantitative Finance', 'Covers research in financial engineering, risk management, and their applications in finance.', NULL);
-- Insert the primary category 'Statistics'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Statistics', 'Covers a broad range of statistical methodologies and their applications.', NULL);
-- Insert the primary category 'Electrical Engineering and Systems Science'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Electrical Engineering and Systems Science', 'Covers a wide range of topics in electrical engineering and systems science, including signal processing, image processing, and control systems.', NULL);
-- Insert the primary category 'Economics'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Economics', 'Covers a wide range of topics in economics including theoretical, empirical, and policy-related research.', NULL);

-- Insert secondary categories under 'Physics'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Astrophysics of Galaxies (astro-ph.GA)', 'This subdivision of astrophysics covers observations, instrumentation, data, surveys, and simulations related to galaxies.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Cosmology and Nongalactic Astrophysics (astro-ph.CO)', 'This section is for papers on the origin, structure, and evolution of the universe, as well as large-scale structures, cosmic microwave background radiation, and other cosmological topics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Earth and Planetary Astrophysics (astro-ph.EP)', 'Covers research on planetary systems and exoplanets, planetary atmospheres, planetary geology, and astrobiology.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'High Energy Astrophysical Phenomena (astro-ph.HE)', 'Includes high-energy phenomena in stars, such as gamma-ray bursts, supernovae, pulsars, and X-ray binaries.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Instrumentation and Methods for Astrophysics (astro-ph.IM)', 'Research on instrumentation, observational techniques, and data analysis methods used in astrophysics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Solar and Stellar Astrophysics (astro-ph.SR)', 'Encompasses research on solar and stellar physics, including stellar evolution, solar dynamics, and activity.', 1);

INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Disordered Systems and Neural Networks (cond-mat.dis-nn)', 'Focuses on disordered systems, complex systems, and properties of neural networks from a condensed matter perspective.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Materials Science (cond-mat.mtrl-sci)', 'Covers research on the properties and behavior of different materials.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Mesoscale and Nanoscale Physics (cond-mat.mes-hall)', 'Research on physical phenomena at the mesoscale and nanoscale.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Other Condensed Matter (cond-mat.other)', 'Includes all condensed matter research not falling into other specific categories.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Quantum Gases (cond-mat.quant-gas)', 'Studies thermal, dynamical, and other properties of quantum gases.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Soft Condensed Matter (cond-mat.soft)', 'Covers polymers, colloids, liquid crystals, and other soft matter.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Statistical Mechanics (cond-mat.stat-mech)', 'Focuses on the statistical properties of matter and underlying theoretical frameworks.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Strongly Correlated Electrons (cond-mat.str-el)', 'Encompasses research on electronic properties of strongly correlated systems.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Superconductivity (cond-mat.supr-con)', 'Studies the phenomena of superconductivity in materials.', 1);

INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'General Relativity and Quantum Cosmology (gr-qc)', 'Covers the broad area of general relativity, quantum gravity, and related subjects.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'High Energy Physics - Experiment (hep-ex)', 'Experimental research in high energy physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'High Energy Physics - Lattice (hep-lat)', 'Research on lattice field theory and lattice gauge theory.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'High Energy Physics - Phenomenology (hep-ph)', 'Studies in the phenomenology of high energy physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'High Energy Physics - Theory (hep-th)', 'Theoretical research in high energy physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Mathematical Physics (math-ph)', 'Theoretical work in the intersection of mathematics and physics.', 1);

INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Adaptation and Self-Organizing Systems (nlin.AO)', 'Studies on self-organization and adaptation in physical systems.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Cellular Automata and Lattice Gases (nlin.CG)', 'Research in cellular automata and lattice gas models.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Chaotic Dynamics (nlin.CD)', 'Focuses on chaotic systems and their behavior.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Pattern Formation and Solitons (nlin.PS)', 'Research on pattern formation and soliton phenomena.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Exactly Solvable and Integrable Systems (nlin.SI)', 'Studies of exactly solvable and integrable systems.', 1);

INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Nuclear Experiment (nucl-ex)', 'Experimental nuclear physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Nuclear Theory (nucl-th)', 'Theoretical nuclear physics.', 1);

INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Accelerator Physics (physics.acc-ph)', 'Research in particle accelerator physics and related technologies.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Atmospheric and Oceanic Physics (physics.ao-ph)', 'Covers atmospheric science and oceanography.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Atomic and Molecular Clusters (physics.atm-clus)', 'Studies in atomic and molecular clusters.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Atomic Physics (physics.atom-ph)', 'Research in atomic physics, including interactions, structural properties, and dynamics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Biological Physics (physics.bio-ph)', 'Research intersecting physics and biology.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Chemical Physics (physics.chem-ph)', 'Interfaces between physics and chemistry.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Classical Physics (physics.class-ph)', 'Fields of classical physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computational Physics (physics.comp-ph)', 'Computational approaches and studies in physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Data Analysis, Statistics and Probability (physics.data-an)', 'Focuses on data analysis, statistical methods, and probabilistic approaches in physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Fluid Dynamics (physics.flu-dyn)', 'Studies in fluid dynamics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'General Physics (physics.gen-ph)', 'General topics in physics not covered elsewhere.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Geophysics (physics.geo-ph)', 'Research in geophysics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'History and Philosophy of Physics (physics.hist-ph)', 'Historical and philosophical aspects of physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Instrumentation and Detectors (physics.ins-det)', 'Development and use of physics instrumentation and detectors.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Medical Physics (physics.med-ph)', 'Studies on applications of physics in medicine.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Optics (physics.optics)', 'Research in optics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Physics Education (physics.ed-ph)', 'Studies focused on education in physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Physics and Society (physics.soc-ph)', 'The intersection of physics with societal impact.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Plasma Physics (physics.plasm-ph)', 'Research in plasma physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Popular Physics (physics.pop-ph)', 'General interest and popular topics in physics.', 1);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Space Physics (physics.space-ph)', 'Studies in space physics.', 1);

INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Quantum Physics (quant-ph)', 'Research on quantum physics.', 1);

-- Insert secondary categories under 'Mathematics'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Algebraic Geometry (math.AG)', 'Research in algebraic geometry.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Algebraic Topology (math.AT)', 'Research in algebraic topology.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Analysis of PDEs (math.AP)', 'Research in the analysis of partial differential equations.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Category Theory (math.CT)', 'Studies in category theory.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Classical Analysis and ODEs (math.CA)', 'Research in classical analysis and ordinary differential equations.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Combinatorics (math.CO)', 'Research in combinatorics.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Commutative Algebra (math.AC)', 'Research in commutative algebra.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Complex Variables (math.CV)', 'Studies in complex analysis, including holomorphic and meromorphic functions.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Differential Geometry (math.DG)', 'Research in differential geometry.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Dynamical Systems (math.DS)', 'Studies in the theory of dynamical systems.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Functional Analysis (math.FA)', 'Research in functional analysis and its applications.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'General Mathematics (math.GM)', 'Broad studies in general mathematics.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'General Topology (math.GN)', 'Research in general topology.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Geometric Topology (math.GT)', 'Studies in geometric topology.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Group Theory (math.GR)', 'Research in group theory and its applications.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'History and Overview (math.HO)', 'Historical studies and general overviews of mathematics.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Information Theory (math.IT)', 'Research in information theory.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'K-Theory and Homology (math.KT)', 'Studies in K-theory and homological algebra.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Logic (math.LO)', 'Research in mathematical logic.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Mathematical Physics (math.MP)', 'The use of mathematical methods in the formulation of physical theories.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Metric Geometry (math.MG)', 'Studies in metric geometry.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Number Theory (math.NT)', 'Research in number theory.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Numerical Analysis (math.NA)', 'Research in numerical methods and their applications.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Operator Algebras (math.OA)', 'Studies in operator algebras and their applications.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Optimization and Control (math.OC)', 'Research in optimization and control theory.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Probability (math.PR)', 'Studies in probability theory.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Quantum Algebra (math.QA)', 'Research in quantum algebras and related structures.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Representation Theory (math.RT)', 'Studies in the representation of algebraic structures.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Rings and Algebras (math.RA)', 'Research in the theory of rings and algebras.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Spectral Theory (math.SP)', 'Research in spectral theory.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Statistics Theory (math.ST)', 'Theoretical studies in statistics.', 2);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Symplectic Geometry (math.SG)', 'Studies in symplectic geometry.', 2);

-- Insert secondary categories under 'Computer Science'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Artificial Intelligence (cs.AI)', 'Covers all areas of AI except Vision, Robotics, Machine Learning, Multiagent Systems, and Computation and Language (Natural Language Processing), which have separate subject areas. In particular, includes Expert Systems, Theorem Proving (although this may overlap with Logic in Computer Science), Knowledge Representation, Planning, and Uncertainty in AI. Roughly includes material in ACM Subject Classes I.2.0, I.2.1, I.2.3, I.2.4, I.2.8, and I.2.11.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computation and Language (cs.CL)', 'Covers natural language processing, with specific application to text, speech, translation, parsing, and machine speech recognition.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computational Complexity (cs.CC)', 'Covers computational complexity theory and related fields including lower bounds, the polynomial hierarchy, communication complexity, probabilistically checkable proofs, circuit complexity, and non-standard complexity theory.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computational Engineering, Finance, and Science (cs.CE)', 'Covers the application of computer science to fields outside of computer science, particularly engineering, computational finance, and the natural sciences.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computational Geometry (cs.CG)', 'Covers computational geometry including algorithmic aspects, geometric optimization, and spatial data structures. Also includes computational topology.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computational and Neural Networks (cs.NE)', 'Covers the engineering and application of artificial neural networks, as well as formulizations of biological neural networks.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computer Science and Game Theory (cs.GT)', 'Covers all aspects of game theory in computer science including algorithmic mechanism design, computational social choice, and price of anarchy.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computer Vision and Pattern Recognition (cs.CV)', 'Covers all aspects of computer vision and pattern recognition including image processing, image analysis, object detection, object recognition, and video analysis.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computers and Society (cs.CY)', 'Covers the impact of computer science on society including studies of ethics within the field, social implications of technology, governmental policy, and the broader impacts of computing on society.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Cryptography and Security (cs.CR)', 'Covers all aspects of cryptography and network security including secure protocols, block ciphers, encryption schemes, and trusted computing.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Data Structures and Algorithms (cs.DS)', 'Covers data structures and algorithms focusing on the theoretical aspects of their design, analysis, and implementation.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Databases (cs.DB)', 'Covers database management systems, distributed database systems, and database design.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Digital Libraries (cs.DL)', 'Covers studies related to digital libraries, which include collection of digital objects, their management, and retrieval mechanisms.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Discrete Mathematics (cs.DM)', 'Covers discrete mathematics including combinatorial mathematics and graph theory.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Distributed, Parallel, and Cluster Computing (cs.DC)', 'Covers distributed systems, parallel computing, and high performance computing including architectures and networking protocols.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Emerging Technologies (cs.ET)', 'Covers novel technologies and recent advances in the field including quantum computing, bioinformatics, and nanotechnology.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Formal Languages and Automata Theory (cs.FL)', 'Covers the study of formal languages, automata, and the theories behind them.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'General Literature (cs.GL)', 'Covers general literature related to computer science, including surveys and other materials not easily categorized into a specific field.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Graphics (cs.GR)', 'Covers all aspects of computer graphics including rendering techniques, visual realism, and hardware acceleration.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Hardware Architecture (cs.AR)', 'Covers computer hardware architectures, including microarchitecture, embedded systems, and digital design.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Human-Computer Interaction (cs.HC)', 'Covers interactive systems and the user interfaces that facilitate this interaction. Includes usability studies and the development of interactive technologies.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Information Retrieval (cs.IR)', 'Covers the theory and practice of information retrieval, including search engines, information filtering, and retrieval models.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Information Theory (cs.IT)', 'Covers the application of information and coding theory to computer science.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Learning (cs.LG)', 'Covers machine learning and related aspects including supervised learning, unsupervised learning, and reinforcement learning.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Logic in Computer Science (cs.LO)', 'Covers logic as applied to computer science including formal verification, automated reasoning, and type theory.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Mathematical Software (cs.MS)', 'Covers software tools used to solve mathematical problems including symbolic computation, numerical analysis, and computational algebra.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Multiagent Systems (cs.MA)', 'Covers systems of decision-making agents, including research on coordination, cooperation, negotiation, communication, and other aspects of multiagent systems.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Multimedia (cs.MM)', 'Covers all aspects of multimedia computing, including audio, video, image processing, and multimedia systems.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Networking and Internet Architecture (cs.NI)', 'Covers networking technologies, protocols, and applications related to computer networks and internet architecture.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Numerical Analysis (cs.NA)', 'Covers numerical methods, numerical linear algebra, and related topics.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Operating Systems (cs.OS)', 'Covers operating systems, resource management, and systems software.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Other Computer Science (cs.OH)', 'Covers areas of computer science not covered by other categories.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Performance (cs.PF)', 'Covers performance analysis and modeling of computational systems.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Programming Languages (cs.PL)', 'Covers all aspects of programming languages including design, implementation, optimization, and formal semantics.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Robotics (cs.RO)', 'Covers all aspects of robotics including control, perception, and interaction.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Social and Information Networks (cs.SI)', 'Covers the study of complex networks and the social and information networks that result from them.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Software Engineering (cs.SE)', 'Covers software engineering methodology, software quality assurance, and formal methods in software engineering.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Sound (cs.SD)', 'Covers all aspects of digital music and sound processing.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Symbolic Computation (cs.SC)', 'Covers symbolic computation including algebraic algorithms, polynomial systems solving, and symbolic integration.', 3);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Systems and Control (cs.SY)', 'Covers systems theory, classical and modern control theory, and systems engineering.', 3);

-- Insert secondary categories under 'Quantitative Biology'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Biomolecules (q-bio.BM)', 'Covers research spanning properties and interactions of biomolecules.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Cell Behavior (q-bio.CB)', 'Covers research on the properties of individual cells and their interactions with the environment, including cell motility, signaling, and morphogenesis.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Genomics (q-bio.GN)', 'Covers the structure, function, evolution, mapping, and editing of genomes.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Molecular Networks (q-bio.MN)', 'Covers the properties and functions of biomolecular networks.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Neurons and Cognition (q-bio.NC)', 'Covers the properties and modeling of neurons, neural circuits, and cognition.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Other Quantitative Biology (q-bio.OT)', 'Covers areas of quantitative biology not covered by other categories.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Populations and Evolution (q-bio.PE)', 'Covers evolutionary dynamics and population biology.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Quantitative Methods (q-bio.QM)', 'Covers the development and application of quantitative analytical methods in biology.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Subcellular Processes (q-bio.SC)', 'Covers research on the dynamics and mechanics of subcellular structures and processes.', 4);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Tissues and Organs (q-bio.TO)', 'Covers the properties, function, and development of tissues and organs.', 4);

-- Insert secondary categories under 'Quantitative Finance'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computational Finance (q-fin.CP)', 'Covers computational techniques and their application in finance.', 5);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Econometrics (q-fin.EC)', 'Covers the use of econometric methods in financial contexts.', 5);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'General Finance (q-fin.GN)', 'Covers general topics in finance and investment.', 5);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Mathematical Finance (q-fin.MF)', 'Covers the mathematical foundations of finance.', 5);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Portfolio Management (q-fin.PM)', 'Covers asset allocation, diversification, and portfolio optimization.', 5);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Pricing of Securities (q-fin.PR)', 'Covers the pricing of derivatives and other financial securities.', 5);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Risk Management (q-fin.RM)', 'Covers strategies and methodologies for managing financial risks.', 5);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Statistical Finance (q-fin.ST)', 'Covers the use of statistical methods in financial applications.', 5);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Trading and Market Microstructure (q-fin.TR)', 'Covers market design, trading, and the impact of market microstructure on asset prices.', 5);

-- Insert secondary categories under 'Statistics'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Applications (stat.AP)', 'Covers statistical applications in various fields.', 6);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Computation (stat.CO)', 'Covers computational techniques in statistical analysis.', 6);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Methodology (stat.ME)', 'Covers the development of statistical methodologies and theoretical aspects.', 6);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Machine Learning (stat.ML)', 'Covers statistical approaches to machine learning.', 6);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Other Statistics (stat.OT)', 'Covers areas of statistics not included in other categories.', 6);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Statistics Theory (stat.TH)', 'Covers the theoretical foundations of statistics.', 6);

-- Insert secondary categories under 'Electrical Engineering and Systems Science'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Audio and Speech Processing (eess.AS)', 'Covers the processing of all aspects of audio and speech signals including analysis, synthesis, recognition, and enhancement.', 7);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Image and Video Processing (eess.IV)', 'Covers the processing and analysis of images and video including techniques like segmentation, compression, and feature extraction.', 7);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Signal Processing (eess.SP)', 'Covers theory and application of signal processing techniques.', 7);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Systems and Control (eess.SY)', 'Covers the principles of systems theory and control systems.', 7);

-- Insert secondary categories under 'Economics'
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Econometrics (econ.EM)', 'Covers the application of statistical techniques to economic data.', 8);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'General Economics (econ.GN)', 'Covers general economic theory and policy.', 8);
INSERT INTO Tag (created_at, updated_at, name, `desc`, parent_id) VALUES (NOW(), NOW(), 'Theoretical Economics (econ.TH)', 'Covers the development and analysis of economic theory.', 8);
