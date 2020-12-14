# RobsTestRepo

This text will be needed to do a first commit :D

Artem:
Ein Paar CI Bsp aus dem GitLab CI Umfeld:
- https://git.t-seclab.telekom.de/seclab/chef/-/blob/master/.gitlab-ci.yml
- das sieht dann so aus: https://git.t-seclab.telekom.de/seclab/chef/-/pipelines/19689
- Und hier im MR sieht Du die Tests (die grünen Häckchen): https://git.t-seclab.telekom.de/seclab/chef/-/merge_requests/533
- https://git.t-seclab.telekom.de/seclab/auri/-/blob/master/.gitlab-ci.yml
- Ergebnis davon: https://git.t-seclab.telekom.de/seclab/auri/-/pipelines/19829
- und hier im MR: https://git.t-seclab.telekom.de/seclab/auri/-/merge_requests/130

GitHub Actions: https://github.com/features/actions
Für den POC reicht es wenn Du eine zweistufige Pipeline damit baust, im ersten Schritt soll so etwas ausgeführt werden:
- echo "First step" && echo "Result from step1" > testfile
Im zweiten Schritt:
- echo "Second step" && cat testfile
Dabei müssen diese Sachen bei Gihub auf PRs reagieren und entsprechend den Status anzeigen, ähnlich wie in unserem GitLab

GitHub Actions: 
https://github.com/features/actions

Was ist CI/CD?
- CI -> Continuous integration
    - CI führt automatisierte build und test phasen aus, damit die Veränderungen zuverlässig in das zentrale repository gemerged werden können
    - Jede neue Integration triggert eine automatische build-and-test Sequent, welche dem Entwicker Feedback gibt und auf Errors hinweist
- CD -> Continuous delivery / Continous deplyment -> hängt davon ab, wie viel das Team code pushed zur Produktion
    - CD ist eine Methode um den Code zu den End-Usern zu bringen
    - erst wenn die CI Phasen abgeschlossen sind, kann man zur CD wechslen 
    - bei kleinen und regelmäßigen software releases ist das Rolbacken im Falle eines Fehlers einfacher
- Continuous Testing
    - automatisiertes Testen 
    - verschiedene Ausprägungen:
        - Unit Testing, prüft individuelle Funktionen auf korrektheit
        - Integration Testing, stellt sicher, dass alle Komponenten und Services zusammen funktionieren 
        - Functional Testing, stellt sicher, dass das neue Feature wie gewollt performed  
        - Acceptance Testing, hier wird Perfomance, Skalierbarkeit, Auslastbarkeit etc. geprüft
        - Static Code analysis, hier wird nach Syntax und Verwundbarkeiten gesucht
        - Automatisierte Tests wie API Tests und Security Testing
- Ist ein Set an practices, das automatisiert Software baut, testet und deployment stufen entwickelt
    - Die Automation hier reduziert Auslieferungszeiten und erhöht Zuverlässigkeit im Entwicklungskreislauf
    - es werden kontinuierlich bugs behoben und features entwickelt 

Was ist eine Pipeline?
- Ist eine ausführbare step-by-step Methode, die jede Software im Entwicklungszyklus durchlaufen muss
- Eine Deployment Pipeline ist der Prozess vom entnehmen des Codes von einer Version Control und dem aufmachen für den User 
- Typischerweise werden neue Versionen einer Applikation um in der Reihenfolge: build, run tests and safe deployed, geführt
- nimmt repetitive Tasks zum Testen ab, es muss nicht mehr manuell getestet werden
- Verschiedene Pipeline Phasen:
    - Commit  Stage: 
        - Ein Quellcode triggert eine Pipeline nach einem Commit
        - Updates get merged with the codebase
        - Es wird Feedback erzeugt
        - CI/CD Tools machen oben aufgeführte Tests und zeigen potenzielle Probleme auf
    - Build Stage:
        - wenn initaial Test vorbei sind
        - wenn notwendig, wird das Programm compiled und es werden durch die Pipeline Container gebildet
        - Die Pipeline verbindet den Code und die Abhängigkeiten um eine ausführbare Software zu generieren
        - Falls hier ein Problem auftritt muss dieses direkt gelöst werden, sonst kann nicht in die nächste Phase gehen
    - Test Stage:
        - Hier werden nochmal Bugs vor dem End-User bewahrt
        - Automatische Tests validieren die Korrektheit des Codes
        - Automatisierte Tests checken das Verhalten des Produkts
        - Die Pipeline gibt Feedback und reported den Status an den neuen Code
    - Deployment Stage:
        - Wenn die Test Stage erfolgreich abgeschlossen ist, geht es hier weiter
        - Die Applikation geht live
        - Mehrere Deployment Umgebungen sind für die End-User reserviert
        - Real-Time überwachung sorgt dafür, dass das neue Feature wie erwartet funktioniert

GitHub Actions - Now with built-in CI/CD! Live from GitHub HQ:
https://www.youtube.com/watch?v=E1OunoCyuhY&feature=emb_rel_end
